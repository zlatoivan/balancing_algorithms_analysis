package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"load_backend/internal/config"
	"load_backend/internal/server"
)

func main() {
	ctx := context.Background()

	err := bootstrap(ctx)
	if err != nil {
		log.Fatalf("[main] bootstrap: %v", err)
	}
}

func bootstrap(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("config.New: %w", err)
	}

	server1 := server.New(cfg.MatrixSize)

	server1.Run(ctx, cfg.HttpPort)
	if err != nil {
		return fmt.Errorf("server1.Run: %w", err)
	}

	<-ctx.Done()

	return nil
}
