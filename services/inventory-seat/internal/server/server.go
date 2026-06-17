// Package server 实现 InventorySeatService（骨架，返回 mock 数据）。
package server

import (
	"context"

	pb "github.com/mioyisama/railway/libs/gen/go/inventoryseat/v1"
)

// Server 实现 pb.InventorySeatServiceServer。
type Server struct {
	pb.UnimplementedInventorySeatServiceServer
}

// New 创建一个 Server。
func New() *Server { return &Server{} }

// LockSeat 返回 mock 锁座结果。
func (s *Server) LockSeat(_ context.Context, _ *pb.LockSeatRequest) (*pb.LockSeatResponse, error) {
	return &pb.LockSeatResponse{LockId: "lock-mock-1", Success: true}, nil
}
