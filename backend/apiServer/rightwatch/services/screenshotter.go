package services

import (
	"fmt"
	"os"

	"github.com/playwright-community/playwright-go"
	"github.com/sirupsen/logrus"
)

// OndiskConfig holds the ondisk.co.kr login info needed to reach gated pages.
// 빈 Username이면 로그인 단계를 건너뛰고 공개 화면만 캡처한다.
type OndiskConfig struct {
	LoginURL string
	Username string
	Password string
}

// CaptureURL launches a headless Chromium, optionally logs into ondisk, navigates
// to pageURL, and returns a full-page JPEG. Pure capture mechanism — no DB access.
// 요청마다 새 브라우저를 띄운다(세션 풀링은 범위 밖). 호출자는 결과 bytes를 파일로 저장한다.
func CaptureURL(pageURL string, cfg OndiskConfig) ([]byte, error) {
	pw, err := playwright.Run()
	if err != nil {
		return nil, fmt.Errorf("playwright run: %w", err)
	}
	defer pw.Stop()

	launchOpts := playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
		Args:     []string{"--no-sandbox", "--disable-dev-shm-usage"},
	}
	if path := os.Getenv("CHROMIUM_PATH"); path != "" {
		launchOpts.ExecutablePath = playwright.String(path)
	}
	browser, err := pw.Chromium.Launch(launchOpts)
	if err != nil {
		return nil, fmt.Errorf("chromium launch: %w", err)
	}
	defer browser.Close()

	ctx, err := browser.NewContext(playwright.BrowserNewContextOptions{
		IgnoreHttpsErrors: playwright.Bool(true),
		UserAgent: playwright.String(
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 " +
				"(KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"),
	})
	if err != nil {
		return nil, fmt.Errorf("new context: %w", err)
	}
	page, err := ctx.NewPage()
	if err != nil {
		return nil, fmt.Errorf("new page: %w", err)
	}

	// 로그인은 best-effort: 실패해도 대상 페이지 캡처는 진행(로그인 월/삭제 화면 자체가 증거).
	if cfg.Username != "" {
		ondiskLogin(page, cfg)
	}

	if _, err := page.Goto(pageURL, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
		Timeout:   playwright.Float(20000),
	}); err != nil {
		// networkidle은 광고/소켓이 많은 페이지에서 타임아웃날 수 있음 — 캡처는 그대로 시도.
		logrus.WithError(err).Warnf("goto %s timed out (capturing current state)", pageURL)
	}

	buf, err := page.Screenshot(playwright.PageScreenshotOptions{
		Type:     playwright.ScreenshotTypeJpeg,
		Quality:  playwright.Int(80),
		FullPage: playwright.Bool(true),
		Timeout:  playwright.Float(30000),
	})
	if err != nil {
		return nil, fmt.Errorf("screenshot: %w", err)
	}
	return buf, nil
}

// ondiskLogin fills the loginFrm (mb_id/mb_pw) and submits via Enter.
// 셀렉터가 안 맞거나 실패해도 조용히 반환 — 호출자가 캡처를 계속한다.
// 폼 구조는 크롤러 스파이더(ondisk_spider3.py)의 loginFrm 로직과 동일.
func ondiskLogin(page playwright.Page, cfg OndiskConfig) {
	if _, err := page.Goto(cfg.LoginURL, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded,
		Timeout:   playwright.Float(20000),
	}); err != nil {
		logrus.WithError(err).Warn("ondisk login page goto failed")
		return
	}
	if err := page.Fill("input[name=\"mb_id\"]", cfg.Username); err != nil {
		logrus.WithError(err).Warn("ondisk login: mb_id fill failed")
		return
	}
	if err := page.Fill("input[name=\"mb_pw\"]", cfg.Password); err != nil {
		logrus.WithError(err).Warn("ondisk login: mb_pw fill failed")
		return
	}
	if err := page.Press("input[name=\"mb_pw\"]", "Enter"); err != nil {
		logrus.WithError(err).Warn("ondisk login submit failed")
		return
	}
	_ = page.WaitForLoadState(playwright.PageWaitForLoadStateOptions{
		State:   playwright.LoadStateNetworkidle,
		Timeout: playwright.Float(15000),
	})
}
