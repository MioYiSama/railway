package grpcserver

import (
	"testing"

	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func TestNewRegistersHealth(t *testing.T) {
	srv := New()
	if srv == nil {
		t.Fatal("New returned nil")
	}
	if _, ok := srv.GetServiceInfo()[healthpb.Health_ServiceDesc.ServiceName]; !ok {
		t.Fatalf("health service not registered, got: %v", srv.GetServiceInfo())
	}
}
