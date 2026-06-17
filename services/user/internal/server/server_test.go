package server

import (
	"context"
	"testing"

	pb "github.com/mioyisama/railway/libs/gen/go/user/v1"
)

func TestGetUser(t *testing.T) {
	resp, err := New().GetUser(context.Background(), &pb.GetUserRequest{Id: "u1"})
	if err != nil {
		t.Fatalf("GetUser error: %v", err)
	}
	if resp.GetId() != "u1" {
		t.Fatalf("Id = %q, want u1", resp.GetId())
	}
}
