package resources

import (
	"context"
	"strings"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestToolUsageGuideResource(t *testing.T) {
	ctx := context.Background()
	req := &mcp.ServerRequest[*mcp.ReadResourceParams]{
		Params: &mcp.ReadResourceParams{
			URI: "tool-usage://guide",
		},
	}

	result, err := ToolUsageGuideResource(ctx, req)
	if err != nil {
		t.Fatalf("ToolUsageGuideResource failed: %v", err)
	}

	if result == nil {
		t.Fatal("Expected non-nil result")
	}

	if len(result.Contents) != 1 {
		t.Fatalf("Expected 1 content item, got %d", len(result.Contents))
	}

	content := result.Contents[0]
	if content.URI != "tool-usage://guide" {
		t.Errorf("Expected URI 'tool-usage://guide', got '%s'", content.URI)
	}

	if content.MIMEType != "text/plain" {
		t.Errorf("Expected MIMEType 'text/plain', got '%s'", content.MIMEType)
	}

	if content.Text == "" {
		t.Error("Expected non-empty text content")
	}

	text := content.Text
	// Verify content contains expected guide information
	expectedStrings := []string{
		"Tool selection guide:",
		"zsv_run(cmd)",
		"zsv_help()",
		"Run zsv count over data.csv",
		`{"cmd":["count","data.csv"]}`,
		"Show zsv help",
		"Validation:",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(text, expected) {
			t.Errorf("Expected text to contain '%s', but it didn't", expected)
		}
	}
}

func TestToolUsageGuideResourceContent(t *testing.T) {
	// Test that the content constant is properly formatted
	if toolUsageGuideContent == "" {
		t.Error("toolUsageGuideContent should not be empty")
	}

	// Verify key sections exist
	if !strings.Contains(toolUsageGuideContent, "Tool selection guide:") {
		t.Error("Content missing 'Tool selection guide:' section")
	}

	if !strings.Contains(toolUsageGuideContent, "Examples:") {
		t.Error("Content missing 'Examples:' section")
	}

	if !strings.Contains(toolUsageGuideContent, "Validation:") {
		t.Error("Content missing 'Validation:' section")
	}
}

func TestToolUsageGuideResourceRequiresRequestParams(t *testing.T) {
	_, err := ToolUsageGuideResource(context.Background(), nil)
	if err == nil {
		t.Fatal("expected error when request is nil")
	}

	_, err = ToolUsageGuideResource(context.Background(), &mcp.ServerRequest[*mcp.ReadResourceParams]{})
	if err == nil {
		t.Fatal("expected error when request params are nil")
	}
}

func TestToolUsageGuideResourceRequiresURI(t *testing.T) {
	_, err := ToolUsageGuideResource(context.Background(), &mcp.ServerRequest[*mcp.ReadResourceParams]{
		Params: &mcp.ReadResourceParams{},
	})
	if err == nil {
		t.Fatal("expected error when uri is missing")
	}
}

func TestRegisterResources(t *testing.T) {
	// Create a test server
	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "test-server",
			Version: "1.0.0",
		},
		&mcp.ServerOptions{
			HasResources: true,
		},
	)

	// This should not panic
	RegisterResources(server)

	// We can't easily test the registration without exposing internals,
	// but we can verify the function doesn't panic
	t.Log("RegisterResources completed without error")
}
