package router

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestPingRoute(t *testing.T) {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	Register(app)

	resp, err := app.Test(httptest.NewRequest("GET", "/api/v1/ping", nil))
	if err != nil {
		t.Fatalf("Test error: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("status = %d, want 200", resp.StatusCode)
	}
}
