package tools

import "github.com/modelcontextprotocol/go-sdk/mcp"

func RegisterTools(server *mcp.Server) {
	if server == nil {
		return
	}

	RegisterRun(server)
	RegisterHelp(server)
}
