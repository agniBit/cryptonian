package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/agniBit/cryptonian/internal/config"
	"github.com/agniBit/cryptonian/internal/logger"
	"github.com/agniBit/cryptonian/pkg/server"
)

func main() {
	cfg := config.LoadConfig()
	logger.Init(cfg)
	logger.InitNewRelic(cfg)

	srv := server.New(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go handleSignals(cancel)

	if err := srv.Start(ctx); err != nil {
		logger.Fatal(ctx, "Server startup failed", err, nil)
	}

	<-ctx.Done()

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Error(context.Background(), "Graceful shutdown failed", err, nil)
	}
}

func handleSignals(cancel context.CancelFunc) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	cancel()
}
