package server

import (
	"context"
	"testing"

	pb "github.com/mioyisama/railway/libs/proto/gen/go/notification/v1"
)

func TestSendNotification(t *testing.T) {
	resp, err := New().SendNotification(context.Background(), &pb.SendNotificationRequest{UserId: "u1", Channel: "sms", Content: "hi"})
	if err != nil {
		t.Fatalf("SendNotification error: %v", err)
	}
	if !resp.GetAccepted() {
		t.Fatal("expected accepted")
	}
}
