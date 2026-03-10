---
name: mcp-tool-usage
description: Tool selection and validation rules for zsv_run and zsv_help in zsv-mcp.
applyTo: ["**/*"]
---

# MCP Tool Usage

Use this guide when deciding which MCP tool to call:

- If user wants to run a zsv command, call `zsv_run(cmd)` with `cmd` as a string array.
- If user wants zsv help text, call `zsv_help()`.

## Examples

- Input intent: "Run zsv count over data.csv"
	Tool: `zsv_run`
	Args: `{"cmd":["count","data.csv"]}`

- Input intent: "Show zsv help"
	Tool: `zsv_help`
	Args: `{}`

## Validation

- Never call `zsv_run` without `cmd`.
- `cmd` must be a non-empty string array.
- `zsv_help` has no required arguments.

## Execution Notes

- `zsv_run` and `zsv_help` use the CLI from `ZSV_PATH` and default to `zsv` when unset.
- Tool results return the captured CLI stdout.

## MCP Context Source

- Resource URI: `tool-usage://guide`
- Prompt name: `tool_usage_guide`

Prefer the resource for background context and use the prompt when explicit examples are requested.
