package services

import (
	"fmt"
	"net/smtp"
	"strings"
)

// SMTPConfig 메일 발송 설정. host가 빈 값이면 dry-run으로 동작한다.
type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

// Enabled SMTP 발송이 설정되어 있는지. host가 비면 dry-run.
func (c SMTPConfig) Enabled() bool {
	return strings.TrimSpace(c.Host) != ""
}

// sender 발신 주소. From이 비면 Username을 사용한다.
func (c SMTPConfig) sender() string {
	if strings.TrimSpace(c.From) != "" {
		return c.From
	}
	return c.Username
}

// BuildNotification CP 탐지 통보 메일의 제목/본문을 한글 템플릿으로 조립한다.
func BuildNotification(cpName, contentTitle, postURL string) (subject, body string) {
	subject = fmt.Sprintf("[저작권 침해 탐지] %s", contentTitle)
	body = fmt.Sprintf(
		"안녕하세요, %s 담당자님.\n\n"+
			"아래 콘텐츠의 불법 유통이 탐지되어 통보드립니다.\n\n"+
			"- 콘텐츠: %s\n"+
			"- 게시물: %s\n\n"+
			"확인 후 삭제 조치를 부탁드립니다.\n",
		cpName, contentTitle, postURL,
	)
	return subject, body
}

// SendMail 메일을 발송한다.
// SMTP 미설정(host 빈 값) 또는 수신자 빈 값이면 dry-run으로 폴백하여 실제 발송하지 않는다.
// 반환: dryRun(실발송 여부), result(결과 메시지), err
func SendMail(cfg SMTPConfig, to, subject, body string) (dryRun bool, result string, err error) {
	if !cfg.Enabled() || strings.TrimSpace(to) == "" {
		return true, "dry-run: SMTP 미설정", nil
	}

	from := cfg.sender()
	msg := buildMessage(from, to, subject, body)
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)

	if err := smtp.SendMail(addr, auth, from, []string{to}, []byte(msg)); err != nil {
		return false, fmt.Sprintf("send failed: %v", err), err
	}
	return false, "sent", nil
}

// buildMessage RFC 5322 형식의 UTF-8 plain 텍스트 메일을 조립한다.
func buildMessage(from, to, subject, body string) string {
	var b strings.Builder
	b.WriteString("From: " + from + "\r\n")
	b.WriteString("To: " + to + "\r\n")
	b.WriteString("Subject: " + subject + "\r\n")
	b.WriteString("MIME-Version: 1.0\r\n")
	b.WriteString("Content-Type: text/plain; charset=UTF-8\r\n")
	b.WriteString("\r\n")
	b.WriteString(strings.ReplaceAll(body, "\n", "\r\n"))
	return b.String()
}
