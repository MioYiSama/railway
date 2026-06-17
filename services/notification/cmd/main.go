package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/mioyisama/railway/libs/config"
	"github.com/mioyisama/railway/libs/grpcserver"
	"github.com/mioyisama/railway/libs/observability"
	pb "github.com/mioyisama/railway/libs/gen/go/notification/v1"
	"github.com/mioyisama/railway/services/notification/internal/server"
)

func main() {
	cfg := config.Load("notification")
	ctx := context.Background()

	shutdown, err := observability.Setup(ctx, cfg.ServiceName, cfg.OTLPEndpoint)
	if err != nil {
		log.Printf("observability setup failed: %v", err)
	} else {
		defer func() { _ = shutdown(ctx) }()
	}

	lis, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		log.Fatalf("listen: %v", err)
	}

	srv := grpcserver.New()
	pb.RegisterNotificationServiceServer(srv, server.New())

	go func() {
		log.Printf("%s gRPC server listening on %s", cfg.ServiceName, cfg.GRPCAddr)
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("serve: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Printf("shutting down %s", cfg.ServiceName)
	srv.GracefulStop()
}
