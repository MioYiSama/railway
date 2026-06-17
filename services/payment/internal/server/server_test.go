package server

import (
	"context"
	"testing"

	pb "github.com/mioyisama/railway/libs/proto/gen/go/payment/v1"
)

func TestCreatePayment(t *testing.T) {
	resp, err := New().CreatePayment(context.Background(), &pb.CreatePaymentRequest{OrderId: "order-1", AmountCents: 55350})
	if err != nil {
		t.Fatalf("CreatePayment error: %v", err)
	}
	if resp.GetStatus() != "SUCCESS" {
		t.Fatalf("Status = %q, want SUCCESS", resp.GetStatus())
	}
}
