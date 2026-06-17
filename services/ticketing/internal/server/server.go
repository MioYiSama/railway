// Package server 实现 TicketingService（骨架，返回 mock 数据）。
package server

import (
	"context"

	pb "github.com/mioyisama/railway/libs/proto/gen/go/ticketing/v1"
)

// Server 实现 pb.TicketingServiceServer。
type Server struct {
	pb.UnimplementedTicketingServiceServer
}

// New 创建一个 Server。
func New() *Server { return &Server{} }

// IssueTicket 返回 mock 电子客票与乘车码。
func (s *Server) IssueTicket(_ context.Context, _ *pb.IssueTicketRequest) (*pb.IssueTicketResponse, error) {
	return &pb.IssueTicketResponse{TicketId: "ticket-mock-1", QrCode: "QR-MOCK"}, nil
}
