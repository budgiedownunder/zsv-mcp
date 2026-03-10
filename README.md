# zsv-mcp

A Go-based Model Context Protocol (MCP) server that exposes tools for the `zsv` CLI (see: https://github.com/liquidaty/zsv).

This project uses the official MCP Go SDK: `github.com/modelcontextprotocol/go-sdk`.

## Features

- `zsv_run`: run `zsv` with explicit argument arrays and return captured stdout.
- `zsv_help`: run `zsv --help` and return captured stdout.
- Configurable runtime values through environment variables.
- Unit tests for core tool behavior.

## Repository Layout

```
.
|- .cursor/
|  |- mcp.json             # Cursor MCP config
|- .github/
|  |- copilot-instructions.md
|- .mcp.json               # Claude Code MCP config
|- .vscode/
|  |- mcp.json             # VS Code / GitHub Copilot MCP config
|  |- settings.json        # VS Code MCP config (legacy format)
|- CLAUDE.md               # Cursor/Claude AI instructions for using zsv MCP tools
|- config/
|  |- config.go            # Environment-based configuration
|- docs/
|  |- TOOL_REFERENCE.md    # detailed tool contracts and examples
|- main.go                 # MCP server entry point
|- main_test.go            # Core handler tests
|- prompts/
|  |- tool_usage.go        # tool usage guidance prompt
|- resources/
|  |- tool_usage.go        # tool usage guidance resource
|- SETUP.md                # VS Code local setup notes
|- test_data/
|  |- data.csv             # sample CSV for tests and tool examples (e.g. zsv sql)
|- tools/
|  |- help.go              # zsv_help tool
|  |- run.go               # zsv_run tool
|  |- run_help_test.go     # zsv_run/zsv_help tool tests
```

## Requirements

- Go 1.23+
- `zsv` CLI available on your `PATH` for `zsv_run` and `zsv_help` tools, or `ZSV_PATH` set to the binary location - see [Installation](https://github.com/liquidaty/zsv?tab=readme-ov-file#installation) details.

## Quick Start

1. Clone and enter the project directory.

2. Install dependencies:

```bash
go mod download
```

3. Build:

```bash
# Linux/macOS
go build -o zsv-mcp

# Windows
go build -o zsv-mcp.exe
```

4. Start the server:

```bash
# Linux/macOS
./zsv-mcp

# Windows
zsv-mcp.exe
```

The server communicates over stdio. See [MCP Client Configuration Files](#mcp-client-configuration-files) to connect it to your editor.

5. Run tests:

```bash
go test ./...
```

## Configuration

Environment variables are read in `config/config.go`.

- `SERVER_NAME` (default: `zsv-mcp`)
- `VERSION` (default: `v1.0.0`)
- `LOG_LEVEL` (default: `info`)
- `ZSV_PATH` (default: `zsv`) used by `zsv_run` and `zsv_help`

Logging notes:

- The server writes logs to `stderr` only to avoid interfering with MCP stdio protocol traffic on `stdout`.
- Supported `LOG_LEVEL` values: `debug`, `info`, `warn` (or `warning`), `error`.
- Unknown `LOG_LEVEL` values fall back to `info` and emit a warning at startup.

PowerShell example:

```powershell
$env:SERVER_NAME = "zsv-mcp"
$env:VERSION = "v1.0.0"
$env:LOG_LEVEL = "debug"
$env:ZSV_PATH = "zsv"
zsv-mcp.exe
```

## MCP Client Configuration Files

The repository ships with pre-configured files so MCP-compatible clients can discover and launch the server without manual setup. All configs expect the binary to be named `zsv-mcp` (no `.exe` extension) — see the cross-platform note in Quick Start above.

### `.mcp.json`

Project-level MCP discovery file read by Claude Code. It registers `zsv-mcp` as a `stdio` server with no extra arguments or environment overrides. The `command` path is relative to the project root, so the client must be launched from that directory (or the path adjusted accordingly).

```json
{
  "mcpServers": {
    "zsv-mcp": {
      "type": "stdio",
      "command": "./zsv-mcp",
      "args": [],
      "env": {
        "ZSV_PATH": "zsv",
        "LOG_LEVEL": "debug"
    }
  }
}
```

### `.vscode/mcp.json`

VS Code-specific MCP server configuration introduced in newer VS Code releases. VS Code reads this file to populate the Chat MCP tool list. It uses `${workspaceFolder}` so the path resolves correctly regardless of where VS Code is opened. GitHub Copilot in VS Code also uses this file.

```json
{
  "servers": {
    "zsv-mcp": {
      "command": "${workspaceFolder}/zsv-mcp",
      "type": "stdio",
      "args": [],
      "env": {
        "ZSV_PATH": "zsv",
        "LOG_LEVEL": "debug"
    }
  }
}
```

### `.vscode/settings.json`

Workspace settings file that registers the server under the `mcpServers` key. Earlier versions of VS Code's MCP integration use this file rather than `.vscode/mcp.json`. The `disabled` field controls whether the server is active, and `alwaysAllow` lists any tool calls that are approved automatically without a user prompt.

```json
{
  "mcpServers": {
    "zsv-mcp": {
      "command": "${workspaceFolder}/zsv-mcp",
      "args": [],
      "disabled": false,
      "alwaysAllow": []
    }
  }
}
```

Both `.vscode/mcp.json` and `.vscode/settings.json` point to the same binary; which file VS Code consults depends on its version, so both are included for compatibility.

### `.cursor/mcp.json`

Cursor MCP configuration. Cursor's support for `${workspaceFolder}` is inconsistent across versions, so this config references the binary by name only. Ensure `zsv-mcp` is on your `PATH` (for example, via `go install`) or replace the command value with the absolute path to the binary.

```json
{
  "mcpServers": {
    "zsv-mcp": {
      "command": "zsv-mcp",
      "args": [],
      "env": {
        "ZSV_PATH": "zsv",
        "LOG_LEVEL": "debug"
    }
  }
}
```

### Claude Desktop

Claude Desktop uses a user-level config file, not a repo-level one. Add the following to your `claude_desktop_config.json`:

- **macOS:** `~/Library/Application Support/Claude/claude_desktop_config.json`
- **Windows:** `%APPDATA%\Claude\claude_desktop_config.json`

```json
{
  "mcpServers": {
    "zsv-mcp": {
      "command": "/absolute/path/to/zsv-mcp",
      "args": [],
      "env": {
        "ZSV_PATH": "zsv",
        "LOG_LEVEL": "debug"
    }
  }
}
```

## Tool Summary

See [TOOL_REFERENCE.md](docs/TOOL_REFERENCE.md) for complete request and response details.

- `zsv_run(cmd[])`
- `zsv_help()`

## Development

- Local editor setup: `SETUP.md`
- Contribution workflow: `CONTRIBUTING.md`
