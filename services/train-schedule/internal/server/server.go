// Package server 实现 TrainScheduleService（骨架，返回 mock 数据）。
package server

import (
	"context"

	pb "github.com/mioyisama/railway/libs/gen/go/trainschedule/v1"
)

// Server 实现 pb.TrainScheduleServiceServer。
type Server struct {
	pb.UnimplementedTrainScheduleServiceServer
}

// New 创建一个 Server。
func New() *Server { return &Server{} }

// GetTrainSchedule 返回 mock 车次时刻信息。
func (s *Server) GetTrainSchedule(_ context.Context, req *pb.GetTrainScheduleRequest) (*pb.GetTrainScheduleResponse, error) {
	return &pb.GetTrainScheduleResponse{
		TrainNo:     req.GetTrainNo(),
		FromStation: "北京南",
		ToStation:   "上海虹桥",
		DepartTime:  "08:00",
	}, nil
}
