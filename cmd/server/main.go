package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/agniBit/cryptonian/internal/config"
	"github.com/agniBit/cryptonian/pkg/server"
)

func main() {
	ctx := context.Background()
	cfg := config.LoadConfig()

	// Channel to listen for interrupt or terminate signals from the OS
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	server.ListenAndServe(ctx, cfg, quit)
}
