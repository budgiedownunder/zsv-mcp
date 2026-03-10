package resources

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const toolUsageGuideContent = `Tool selection guide:
- If user wants to run a zsv command -> call zsv_run(cmd)
- If user wants zsv help text -> call zsv_help()

Examples:
1) Input intent: "Run zsv count over data.csv"
	Tool: zsv_run
	Args: {"cmd":["count","data.csv"]}

2) Input intent: "Run zsv sql with inline SQL"
	Tool: zsv_run
	Args: {"cmd":["sql","test.csv","select avg(cast([age] as real)) from data"]}

3) Input intent: "Show zsv help"
	Tool: zsv_help
	Args: {}

Validation:
- Never call zsv_run without "cmd"
- cmd must be a non-empty string array
- zsv_help has no required arguments`

// RegisterResources registers resources that provide context and documentation.
func RegisterResources(server *mcp.Server) {
	if server == nil {
		return
	}

	server.AddResource(&mcp.Resource{
		URI:         "tool-usage://guide",
		Name:        "Tool Usage Guide",
		Description: "Guidelines for using zsv_run and zsv_help tools correctly",
		MIMEType:    "text/plain",
	}, ToolUsageGuideResource)
}

// ToolUsageGuideResource returns the tool usage guide content as a resource.
func ToolUsageGuideResource(_ context.Context, req *mcp.ServerRequest[*mcp.ReadResourceParams]) (*mcp.ReadResourceResult, error) {
	if req == nil || req.Params == nil {
		return nil, fmt.Errorf("resource request params are required")
	}
	if req.Params.URI == "" {
		return nil, fmt.Errorf("resource URI is required")
	}

	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{{
			URI:      req.Params.URI,
			MIMEType: "text/plain",
			Text:     toolUsageGuideContent,
		}},
	}, nil
}
