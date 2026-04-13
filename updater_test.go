package main

import (
	"errors"
	"testing"
	"time"
)

func TestFormatSuccessNotification(t *testing.T) {
	notAfter := time.Date(2026, 5, 1, 2, 3, 4, 0, time.UTC)

	got := formatSuccessNotification("example.com", "cert-123", notAfter)
	want := "**域名**: example.com\n**到期时间**: 2026-05-01T02:03:04Z\n**证书ID**: cert-123"
	if got != want {
		t.Fatalf("formatSuccessNotification() = %q, want %q", got, want)
	}
}

func TestFormatFailureNotification(t *testing.T) {
	got := formatFailureNotification("example.com", "verify_external", errors.New("fingerprint mismatch"))
	want := "**域名**: example.com\n**阶段**: verify_external\n**错误**: fingerprint mismatch"
	if got != want {
		t.Fatalf("formatFailureNotification() = %q, want %q", got, want)
	}
}
