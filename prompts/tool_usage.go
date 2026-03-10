package prompts

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// RegisterPrompts registers prompt templates that help clients choose and call tools.
func RegisterPrompts(server *mcp.Server) {
	if server == nil {
		return
	}

	server.AddPrompt(&mcp.Prompt{
		Name:        "tool_usage_guide",
		Description: "How to choose and call zsv_run and zsv_help correctly",
	}, ToolUsageGuidePrompt)
}

// ToolUsageGuidePrompt returns a compact guide with canonical examples.
func ToolUsageGuidePrompt(_ context.Context, _ *mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	return &mcp.GetPromptResult{
		Description: "Guidance for selecting and calling tools",
		Messages: []*mcp.PromptMessage{
			{
				Role: "user",
				Content: &mcp.TextContent{
					Text: `Tool selection guide:
- If user wants to run a zsv command -> call zsv_run(cmd)
- If user wants zsv help text -> call zsv_help()

Examples:
1) Input intent: "Run zsv count over data.csv"
	Tool: zsv_run
	Args: {"cmd":["count","data.csv"]}

1) Input intent: "Run zsv sql with inline SQL"
	Tool: zsv_run
	Args: {"cmd":["sql","test.csv","select avg(cast([age] as real)) from data"]}

2) Input intent: "Show zsv help"
	Tool: zsv_help
	Args: {}

Validation:
- Never call zsv_run without "cmd"
- cmd must be a non-empty string array
- zsv_help has no required arguments`,
				},
			},
		},
	}, nil
}
