// Package server 实现 OrderService（骨架，返回 mock 数据）。
package server

import (
	"context"

	pb "github.com/mioyisama/railway/libs/gen/go/order/v1"
)

// Server 实现 pb.OrderServiceServer。
type Server struct {
	pb.UnimplementedOrderServiceServer
}

// New 创建一个 Server。
func New() *Server { return &Server{} }

// CreateOrder 返回 mock 订单。
func (s *Server) CreateOrder(_ context.Context, _ *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	return &pb.CreateOrderResponse{OrderId: "order-mock-1", Status: "PENDING_PAYMENT"}, nil
}
