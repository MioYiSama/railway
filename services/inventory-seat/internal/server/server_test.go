package server

import (
	"context"
	"testing"

	pb "github.com/mioyisama/railway/libs/proto/gen/go/inventoryseat/v1"
)

func TestLockSeat(t *testing.T) {
	resp, err := New().LockSeat(context.Background(), &pb.LockSeatRequest{TrainNo: "G1", SeatType: "二等座", Quantity: 1})
	if err != nil {
		t.Fatalf("LockSeat error: %v", err)
	}
	if !resp.GetSuccess() {
		t.Fatal("expected success")
	}
}
