package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
	"github.com/mioyisama/railway/libs/config"
	"github.com/mioyisama/railway/libs/httpserver"
	"github.com/mioyisama/railway/libs/observability"
	"github.com/mioyisama/railway/services/api-gateway/internal/router"
)

func main() {
	cfg := config.Load("api-gateway")
	ctx := context.Background()

	shutdown, err := observability.Setup(ctx, cfg.ServiceName, cfg.OTLPEndpoint)
	if err != nil {
		log.Printf("observability setup failed: %v", err)
	} else {
		defer func() { _ = shutdown(ctx) }()
	}

	app := httpserver.New(cfg.ServiceName)
	router.Register(app)

	go func() {
		log.Printf("api-gateway HTTP server listening on %s", cfg.HTTPAddr)
		if err := app.Listen(cfg.HTTPAddr, fiber.ListenConfig{DisableStartupMessage: true}); err != nil {
			log.Fatalf("listen: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Printf("shutting down api-gateway")
	_ = app.Shutdown()
}
