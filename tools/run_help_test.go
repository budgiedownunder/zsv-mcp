package tools

import (
	"context"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestRunHandlerRequiresCmd(t *testing.T) {
	_, _, err := RunHandler(context.Background(), nil, RunInput{})
	if err == nil {
		t.Fatal("expected error when cmd is missing")
	}
}

func TestRunHandlerRejectsMultiLineCmdArg(t *testing.T) {
	_, _, err := RunHandler(context.Background(), nil, RunInput{Cmd: []string{"count", "data.csv\nhelp"}})
	if err == nil {
		t.Fatal("expected error for multi-line cmd")
	}

	if !strings.Contains(err.Error(), "single line") {
		t.Fatalf("expected single-line error, got %q", err.Error())
	}
}

func TestRunHandlerRequiresFirstArg(t *testing.T) {
	_, _, err := RunHandler(context.Background(), nil, RunInput{Cmd: []string{"", "data.csv"}})
	if err == nil {
		t.Fatal("expected error when first arg is empty")
	}

	if !strings.Contains(err.Error(), "cmd[0] is required") {
		t.Fatalf("expected first-arg required error, got %q", err.Error())
	}
}

func TestRunHandlerCapturesStdout(t *testing.T) {
	restore := setCommandContextForTest(t)
	defer restore()

	t.Setenv("ZSV_PATH", "zsv-custom")

	_, output, err := RunHandler(context.Background(), nil, RunInput{Cmd: []string{"count", "data.csv"}})
	if err != nil {
		t.Fatalf("RunHandler failed: %v", err)
	}

	if output.Status != "success" {
		t.Fatalf("expected success status, got %q", output.Status)
	}

	if got := strings.Join(output.Command, " "); got != "count data.csv" {
		t.Fatalf("unexpected command value %q", got)
	}

	if got := output.Stdout; got != "zsv-custom|count|data.csv" {
		t.Fatalf("unexpected stdout: %q", got)
	}
}

func TestRunHandlerStripsTrailingLineFeeds(t *testing.T) {
	restore := setCommandContextForTest(t)
	defer restore()

	_, output, err := RunHandler(context.Background(), nil, RunInput{Cmd: []string{"count", "data.csv"}})
	if err != nil {
		t.Fatalf("RunHandler failed: %v", err)
	}

	if strings.HasSuffix(output.Stdout, "\n") || strings.HasSuffix(output.Stdout, "\r") {
		t.Fatalf("expected stdout without trailing line-feed, got %q", output.Stdout)
	}
}

func TestHelpHandlerRunsDashHelp(t *testing.T) {
	restore := setCommandContextForTest(t)
	defer restore()

	t.Setenv("ZSV_PATH", "zsv")

	_, output, err := HelpHandler(context.Background(), nil, struct{}{})
	if err != nil {
		t.Fatalf("HelpHandler failed: %v", err)
	}

	if output.Status != "success" {
		t.Fatalf("expected success status, got %q", output.Status)
	}

	if got := strings.TrimSpace(output.Stdout); got != "zsv|--help" {
		t.Fatalf("unexpected stdout: %q", got)
	}
}

func TestRunHandlerReturnsCommandErrorText(t *testing.T) {
	restore := setCommandContextForTest(t)
	defer restore()

	t.Setenv("ZSV_PATH", "zsv")

	_, _, err := RunHandler(context.Background(), nil, RunInput{Cmd: []string{"fail", "now"}})
	if err == nil {
		t.Fatal("expected command failure")
	}

	msg := err.Error()
	if !strings.Contains(msg, "failed") {
		t.Fatalf("expected failure marker in error, got %q", msg)
	}
	if !strings.Contains(msg, "simulated stderr") {
		t.Fatalf("expected stderr text in error, got %q", msg)
	}
	if !strings.Contains(msg, "run command fail now") {
		t.Fatalf("expected wrapped command context in error, got %q", msg)
	}
}

func setCommandContextForTest(t *testing.T) func() {
	t.Helper()
	original := commandContext
	commandContext = func(ctx context.Context, name string, args ...string) *exec.Cmd {
		helperArgs := []string{"-test.run=TestHelperProcess", "--", name}
		helperArgs = append(helperArgs, args...)
		cmd := exec.CommandContext(ctx, os.Args[0], helperArgs...)
		cmd.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")
		return cmd
	}

	return func() {
		commandContext = original
	}
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	args := os.Args
	split := 0
	for i, value := range args {
		if value == "--" {
			split = i
			break
		}
	}

	if split == 0 || split+1 >= len(args) {
		_, _ = os.Stderr.WriteString("missing helper args\n")
		os.Exit(2)
	}

	target := args[split+1]
	toolArgs := args[split+2:]

	if len(toolArgs) > 0 && toolArgs[0] == "fail" {
		_, _ = os.Stderr.WriteString("simulated stderr\n")
		os.Exit(1)
	}

	line := target
	if len(toolArgs) > 0 {
		line += "|" + strings.Join(toolArgs, "|")
	}

	_, _ = os.Stdout.WriteString(line + "\n")
	os.Exit(0)
}

func TestRegisterRunAndHelp(t *testing.T) {
	server := mcp.NewServer(
		&mcp.Implementation{Name: "test", Version: "1.0.0"},
		&mcp.ServerOptions{HasTools: true},
	)

	RegisterRun(server)
	RegisterHelp(server)
}
