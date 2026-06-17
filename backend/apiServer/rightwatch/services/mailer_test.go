package services

import (
	"strings"
	"testing"
)

func TestSMTPConfig_Enabled(t *testing.T) {
	if (SMTPConfig{Host: ""}).Enabled() {
		t.Error("empty host should be disabled (dry-run)")
	}
	if (SMTPConfig{Host: "   "}).Enabled() {
		t.Error("whitespace host should be disabled (dry-run)")
	}
	if !(SMTPConfig{Host: "smtp.example.com"}).Enabled() {
		t.Error("non-empty host should be enabled")
	}
}

func TestSendMail_DryRunWhenHostEmpty(t *testing.T) {
	cfg := SMTPConfig{Host: ""}
	dryRun, result, err := SendMail(cfg, "cp@example.com", "subj", "body")
	if err != nil {
		t.Fatalf("dry-run must not error: %v", err)
	}
	if !dryRun {
		t.Error("empty host should force dry-run")
	}
	if !strings.Contains(result, "dry-run") {
		t.Errorf("result should mention dry-run, got %q", result)
	}
}

func TestSendMail_DryRunWhenToEmpty(t *testing.T) {
	cfg := SMTPConfig{Host: "smtp.example.com", Port: 587}
	dryRun, _, err := SendMail(cfg, "", "subj", "body")
	if err != nil {
		t.Fatalf("dry-run must not error: %v", err)
	}
	if !dryRun {
		t.Error("empty recipient should force dry-run")
	}
}

func TestBuildNotification(t *testing.T) {
	subject, body := BuildNotification("ABC미디어", "테스트영화", "https://webhard.example/post/1")
	if !strings.Contains(subject, "테스트영화") {
		t.Errorf("subject should contain content title, got %q", subject)
	}
	if !strings.Contains(body, "ABC미디어") {
		t.Errorf("body should contain cp name, got %q", body)
	}
	if !strings.Contains(body, "https://webhard.example/post/1") {
		t.Errorf("body should contain post URL, got %q", body)
	}
}

func TestSMTPConfig_SenderFallback(t *testing.T) {
	if got := (SMTPConfig{Username: "u@x.com"}).sender(); got != "u@x.com" {
		t.Errorf("empty From should fall back to Username, got %q", got)
	}
	if got := (SMTPConfig{Username: "u@x.com", From: "f@x.com"}).sender(); got != "f@x.com" {
		t.Errorf("From should take precedence, got %q", got)
	}
}
