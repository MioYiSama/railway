// Package server 实现 NotificationService（骨架，全部 mock）。
package server

import (
	"context"

	pb "github.com/mioyisama/railway/libs/gen/go/notification/v1"
)

// Server 实现 pb.NotificationServiceServer。
type Server struct {
	pb.UnimplementedNotificationServiceServer
}

// New 创建一个 Server。
func New() *Server { return &Server{} }

// SendNotification 返回 mock 受理结果。
func (s *Server) SendNotification(_ context.Context, _ *pb.SendNotificationRequest) (*pb.SendNotificationResponse, error) {
	return &pb.SendNotificationResponse{NotificationId: "notif-mock-1", Accepted: true}, nil
}
