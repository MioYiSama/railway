package httpserver

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthz(t *testing.T) {
	app := New("test-service")
	req := httptest.NewRequest("GET", "/healthz", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Test error: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("status = %d, want 200", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	if !strings.Contains(string(body), "\"status\":\"ok\"") {
		t.Fatalf("body = %s, want status ok", body)
	}
}
