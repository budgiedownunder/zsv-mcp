package tools

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type HelpOutput struct {
	Stdout string `json:"stdout"`
	Status string `json:"status"`
}

func HelpHandler(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (
	*mcp.CallToolResult,
	HelpOutput,
	error,
) {
	slog.Debug("handling zsv_help")
	stdout, err := runZSV(ctx, []string{"--help"})
	if err != nil {
		slog.Error("zsv_help failed", "error", err)
		return nil, HelpOutput{}, fmt.Errorf("run help command: %w", err)
	}

	slog.Info("zsv_help succeeded")

	return nil, HelpOutput{
		Stdout: stdout,
		Status: "success",
	}, nil
}

func RegisterHelp(server *mcp.Server) {
	if server == nil {
		return
	}

	mcp.AddTool(server,
		&mcp.Tool{
			Name:        "zsv_help",
			Description: "Runs zsv --help and returns captured stdout",
		},
		HelpHandler,
	)
}
