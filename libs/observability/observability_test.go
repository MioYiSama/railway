package observability

import (
	"context"
	"testing"
	"time"
)

func TestSetupShutdown(t *testing.T) {
	ctx := context.Background()
	shutdown, err := Setup(ctx, "test-service", "localhost:4317")
	if err != nil {
		t.Fatalf("Setup error: %v", err)
	}
	if shutdown == nil {
		t.Fatal("shutdown func is nil")
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := shutdown(ctx); err != nil {
		t.Fatalf("shutdown error: %v", err)
	}
}
