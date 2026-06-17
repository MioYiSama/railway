// Package server 实现 PaymentService（骨架，全部 mock，不接真实渠道）。
package server

import (
	"context"

	pb "github.com/mioyisama/railway/libs/proto/gen/go/payment/v1"
)

// Server 实现 pb.PaymentServiceServer。
type Server struct {
	pb.UnimplementedPaymentServiceServer
}

// New 创建一个 Server。
func New() *Server { return &Server{} }

// CreatePayment 返回 mock 支付单（直接置为成功）。
func (s *Server) CreatePayment(_ context.Context, _ *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	return &pb.CreatePaymentResponse{PaymentId: "pay-mock-1", Status: "SUCCESS"}, nil
}
