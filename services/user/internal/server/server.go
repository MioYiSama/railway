// Package server 实现 UserService（骨架，返回 mock 数据）。
package server

import (
	"context"

	pb "github.com/mioyisama/railway/libs/gen/go/user/v1"
)

// Server 实现 pb.UserServiceServer。
type Server struct {
	pb.UnimplementedUserServiceServer
}

// New 创建一个 Server。
func New() *Server { return &Server{} }

// GetUser 返回 mock 用户信息。
func (s *Server) GetUser(_ context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{Id: req.GetId(), Name: "mock-user", Verified: true}, nil
}
