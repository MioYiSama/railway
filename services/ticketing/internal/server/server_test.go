package server

import (
	"context"
	"testing"

	pb "github.com/mioyisama/railway/libs/gen/go/ticketing/v1"
)

func TestIssueTicket(t *testing.T) {
	resp, err := New().IssueTicket(context.Background(), &pb.IssueTicketRequest{OrderId: "order-1"})
	if err != nil {
		t.Fatalf("IssueTicket error: %v", err)
	}
	if resp.GetTicketId() == "" {
		t.Fatal("expected non-empty ticket id")
	}
}
