package server

import (
	"context"
	"testing"

	pb "github.com/mioyisama/railway/libs/proto/gen/go/order/v1"
)

func TestCreateOrder(t *testing.T) {
	resp, err := New().CreateOrder(context.Background(), &pb.CreateOrderRequest{UserId: "u1", TrainNo: "G1", SeatType: "二等座"})
	if err != nil {
		t.Fatalf("CreateOrder error: %v", err)
	}
	if resp.GetOrderId() == "" {
		t.Fatal("expected non-empty order id")
	}
}
