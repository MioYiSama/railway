// Package server 实现 TicketQueryService（骨架，返回 mock 数据）。
package server

import (
	"context"

	pb "github.com/mioyisama/railway/libs/gen/go/ticketquery/v1"
)

// Server 实现 pb.TicketQueryServiceServer。
type Server struct {
	pb.UnimplementedTicketQueryServiceServer
}

// New 创建一个 Server。
func New() *Server { return &Server{} }

// QueryTickets 返回 mock 余票列表。
func (s *Server) QueryTickets(_ context.Context, _ *pb.QueryTicketsRequest) (*pb.QueryTicketsResponse, error) {
	return &pb.QueryTicketsResponse{
		Tickets: []*pb.Ticket{
			{TrainNo: "G1", SeatType: "二等座", Remaining: 42, PriceCents: 55350},
		},
	}, nil
}
