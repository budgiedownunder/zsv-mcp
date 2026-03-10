package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/budgiedownunder/zsv-mcp/config"
	"github.com/budgiedownunder/zsv-mcp/prompts"
	"github.com/budgiedownunder/zsv-mcp/resources"
	"github.com/budgiedownunder/zsv-mcp/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	cfg := config.Load()
	level, knownLevel := config.ParseLogLevel(cfg.LogLevel)
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(logger)

	if !knownLevel {
		slog.Warn("unknown LOG_LEVEL; falling back to info", "log_level", cfg.LogLevel)
	}

	slog.Info("starting mcp server", "server_name", cfg.ServerName, "version", cfg.Version, "log_level", level.String())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sigCh)
	go func() {
		<-sigCh
		slog.Info("received shutdown signal")
		cancel()
	}()

	// Create server
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    cfg.ServerName,
			Version: cfg.Version,
		},
		&mcp.ServerOptions{
			HasTools:     true,
			HasResources: true,
			HasPrompts:   true,
		},
	)

	// Register components
	tools.RegisterTools(server)
	resources.RegisterResources(server)
	prompts.RegisterPrompts(server)

	// Run server
	transport := &mcp.StdioTransport{}
	if err := server.Run(ctx, transport); err != nil {
		if errors.Is(err, context.Canceled) {
			slog.Info("server stopped after context cancellation")
			return
		}
		slog.Error("server error", "error", err)
		os.Exit(1)
	}

	slog.Info("server stopped")
}
