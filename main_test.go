package main

import (
	"context"
	"strings"
	"testing"

	"github.com/budgiedownunder/zsv-mcp/prompts"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestToolUsageGuidePromptContainsExamples(t *testing.T) {
	res, err := prompts.ToolUsageGuidePrompt(context.Background(), &mcp.GetPromptRequest{})
	if err != nil {
		t.Fatalf("ToolUsageGuidePrompt failed: %v", err)
	}

	if len(res.Messages) == 0 {
		t.Fatal("expected at least one prompt message")
	}

	text, ok := res.Messages[0].Content.(*mcp.TextContent)
	if !ok {
		t.Fatalf("expected text content, got %T", res.Messages[0].Content)
	}

	if !strings.Contains(text.Text, "zsv_run") {
		t.Error("prompt missing zsv_run usage example")
	}

	if !strings.Contains(text.Text, "zsv_help") {
		t.Error("prompt missing zsv_help usage example")
	}

	if !strings.Contains(text.Text, `{"cmd":["count","data.csv"]}`) {
		t.Error("prompt missing run args example")
	}
}
