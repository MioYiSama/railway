package server

import (
	"context"
	"testing"

	pb "github.com/mioyisama/railway/libs/gen/go/ticketquery/v1"
)

func TestQueryTickets(t *testing.T) {
	resp, err := New().QueryTickets(context.Background(), &pb.QueryTicketsRequest{FromStation: "BJ", ToStation: "SH"})
	if err != nil {
		t.Fatalf("QueryTickets error: %v", err)
	}
	if len(resp.GetTickets()) == 0 {
		t.Fatal("expected at least one ticket")
	}
}
