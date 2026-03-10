# Local VS Code Setup

This guide explains how to run `zsv-mcp` from VS Code Chat for local testing.

For full project documentation, see:

- [README.md](./README.md)
- [TOOL_REFERENCE.md](./docs/TOOL_REFERENCE.md)
- [CONTRIBUTING.md](./CONTRIBUTING.md)

## Quick Start

1. Build the server binary:

```bash
go build -o zsv-mcp
```

> Using `-o zsv-mcp` (without `.exe`) produces a consistently named binary on Windows, Linux, and macOS. The workspace config files all expect this name.

2. Confirm your workspace MCP settings point to this binary (typically in `.vscode/settings.json`).

3. Reload VS Code:

- Command Palette -> `Developer: Reload Window`

4. Open Chat and verify `zsv-mcp` tools are available.

## Validate Manually

Run the server directly:

```bash
./zsv-mcp
```

Expected behavior:

- The process starts and waits for stdio MCP messages.
- It appears idle until a client connects.

## Environment Variables

PowerShell example:

```powershell
$env:SERVER_NAME = "zsv-mcp"
$env:VERSION = "v1.0.0"
$env:LOG_LEVEL = "debug"
$env:ZSV_PATH = "zsv"
zsv-mcp.exe
```

`LOG_LEVEL` supports `debug`, `info`, `warn`/`warning`, and `error`.

The server logs to `stderr` so MCP `stdout` traffic stays protocol-safe.

## Troubleshooting

If tools do not appear in Chat:

1. Open `View -> Output` and select `MCP Server`.
2. Check `.vscode/settings.json` for syntax errors.
3. Rebuild the binary and reload VS Code.

If `zsv_run` or `zsv_help` fails:

1. Verify `zsv` is installed and callable from terminal.
2. Set `ZSV_PATH` explicitly to the correct executable path.
3. For `zsv_run`, provide `cmd` as an array of args (for example: `{"cmd":["count","data.csv"]}`).
4. Inspect MCP server output for command error details.


