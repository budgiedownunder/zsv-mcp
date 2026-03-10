package tools

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const defaultZSVPath = "zsv"

var commandContext = exec.CommandContext

type RunInput struct {
	Cmd []string `json:"cmd" jsonschema:"required"`
}

type RunOutput struct {
	Command []string `json:"command"`
	Stdout  string   `json:"stdout"`
	Status  string   `json:"status"`
}

func zsvPath() string {
	if value := os.Getenv("ZSV_PATH"); value != "" {
		return value
	}

	return defaultZSVPath
}

func parseCmdArg(cmd []string) ([]string, error) {
	if len(cmd) == 0 {
		return nil, fmt.Errorf("cmd is required")
	}

	args := make([]string, len(cmd))
	for i, value := range cmd {
		if strings.ContainsAny(value, "\r\n") {
			return nil, fmt.Errorf("cmd[%d] must be a single line", i)
		}

		if i == 0 && strings.TrimSpace(value) == "" {
			return nil, fmt.Errorf("cmd[0] is required")
		}

		args[i] = value
	}

	return args, nil
}

func runZSV(ctx context.Context, args []string) (string, error) {
	if ctx.Err() != nil {
		return "", ctx.Err()
	}

	binary := zsvPath()
	slog.Debug("running zsv command", "binary", binary, "arg_count", len(args), "subcommand", args[0])
	command := commandContext(ctx, binary, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &stdout
	command.Stderr = &stderr

	if err := command.Run(); err != nil {
		errText := strings.TrimSpace(stderr.String())
		slog.Error("zsv command failed", "binary", binary, "subcommand", args[0], "error", err, "has_stderr", errText != "")
		if errText != "" {
			return "", fmt.Errorf("%s %s failed: %w: %s", binary, strings.Join(args, " "), err, errText)
		}

		return "", fmt.Errorf("%s %s failed: %w", binary, strings.Join(args, " "), err)
	}

	slog.Debug("zsv command completed", "binary", binary, "subcommand", args[0], "stdout_bytes", stdout.Len())

	// Normalize CLI output by removing terminal line endings only.
	return strings.TrimRight(stdout.String(), "\r\n"), nil
}

func RunHandler(ctx context.Context, _ *mcp.CallToolRequest, input RunInput) (
	*mcp.CallToolResult,
	RunOutput,
	error,
) {
	args, err := parseCmdArg(input.Cmd)
	if err != nil {
		slog.Warn("invalid zsv_run input", "error", err)
		return nil, RunOutput{}, err
	}

	stdout, err := runZSV(ctx, args)
	if err != nil {
		slog.Error("zsv_run failed", "subcommand", args[0], "error", err)
		return nil, RunOutput{}, fmt.Errorf("run command %s: %w", strings.Join(input.Cmd, " "), err)
	}

	slog.Info("zsv_run succeeded", "subcommand", args[0])

	return nil, RunOutput{
		Command: append([]string(nil), input.Cmd...),
		Stdout:  stdout,
		Status:  "success",
	}, nil
}

func RegisterRun(server *mcp.Server) {
	if server == nil {
		return
	}

	mcp.AddTool(server,
		&mcp.Tool{
			Name:        "zsv_run",
			Description: "Runs zsv with the provided cmd argument and returns captured stdout",
		},
		RunHandler,
	)
}
