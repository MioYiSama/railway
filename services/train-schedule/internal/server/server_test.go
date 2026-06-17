package server

import (
	"context"
	"testing"

	pb "github.com/mioyisama/railway/libs/proto/gen/go/trainschedule/v1"
)

func TestGetTrainSchedule(t *testing.T) {
	resp, err := New().GetTrainSchedule(context.Background(), &pb.GetTrainScheduleRequest{TrainNo: "G1"})
	if err != nil {
		t.Fatalf("GetTrainSchedule error: %v", err)
	}
	if resp.GetTrainNo() != "G1" {
		t.Fatalf("TrainNo = %q, want G1", resp.GetTrainNo())
	}
}
