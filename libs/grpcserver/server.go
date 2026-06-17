// Package grpcserver 提供带健康检查与反射的 gRPC server 引导封装（骨架）。
package grpcserver

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

// New 创建一个已注册健康检查与反射服务的 gRPC server。
// 业务服务在返回的 server 上注册自己的实现后再 Serve。
func New(opts ...grpc.ServerOption) *grpc.Server {
	srv := grpc.NewServer(opts...)

	hs := health.NewServer()
	hs.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	healthpb.RegisterHealthServer(srv, hs)

	reflection.Register(srv)
	return srv
}
