package config

import "testing"

func TestLoadDefaults(t *testing.T) {
	t.Setenv("GRPC_ADDR", "")
	cfg := Load("user")
	if cfg.ServiceName != "user" {
		t.Fatalf("ServiceName = %q, want user", cfg.ServiceName)
	}
	if cfg.GRPCAddr != ":50051" {
		t.Fatalf("GRPCAddr = %q, want :50051", cfg.GRPCAddr)
	}
}

func TestGetenvOverride(t *testing.T) {
	t.Setenv("HTTP_ADDR", ":9000")
	if got := Getenv("HTTP_ADDR", ":8080"); got != ":9000" {
		t.Fatalf("Getenv = %q, want :9000", got)
	}
}
