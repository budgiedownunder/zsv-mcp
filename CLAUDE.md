# Claude Instructions

When using MCP tools from `zsv-mcp`, follow these rules:

- If the user wants to run a zsv command, call `zsv_run(cmd)` where `cmd` is a string array.
- If the user wants zsv help text, call `zsv_help()`.
- Always provide required arguments.
- Never call `zsv_run` without `cmd`.
- `cmd` must be a non-empty string array.
- `zsv_help` has no required arguments.
- Use `zsv_help` to learn about supported zsv arguments.

## Execution Notes

- `zsv_run` and `zsv_help` execute the CLI from `ZSV_PATH` (default: `zsv`).
- Tool results return captured command stdout from the CLI.

## MCP Context

ALWAYS read the `tool-usage://guide` resource before running any zsv command.
It contains usage examples that must be followed.

- Resource: `tool-usage://guide`
- Prompt: `tool_usage_guide` (use when explicit user-facing examples are needed)

## SQL Command Notes

- When using `zsv sql`, CSV files are loaded as tables named `data`, `data2`, `data3`, ...
- Cast numeric columns explicitly, e.g. `cast([age] as real)`
