# Tool Reference

This document describes the MCP tools registered in `tools/registry.go`.

## `zsv_run`

Runs `zsv` using an explicit argument array.

Input:

```json
{
  "cmd": ["count", "data.csv"]
}
```

Input fields:

- `cmd` (string array, required). The first item is the `zsv` subcommand.

Success output:

```json
{
  "command": ["count", "data.csv"],
  "stdout": "<captured zsv stdout>",
  "status": "success"
}
```

Behavior notes:

- Uses `ZSV_PATH` when set; otherwise defaults to `zsv`.
- Passes args directly to `exec.CommandContext` without shell parsing.
- Returns stdout with trailing `\r` and `\n` removed.
- On command failure, includes stderr text in the error when available.

Validation:

- Returns an error if `cmd` is empty.
- Returns an error if `cmd[0]` is blank.
- Returns an error if any arg contains a newline.

## `zsv_help`

Runs `zsv --help` and returns captured stdout.

Input:

```json
{}
```

Success output:

```json
{
  "stdout": "<captured zsv help text>",
  "status": "success"
}
```

Behavior notes:

- Uses the same `ZSV_PATH` resolution and execution path as `zsv_run`.

## Resource and Prompt Context

The repository also includes helper context definitions:

- Resource URI: `tool-usage://guide` in `resources/tool_usage.go`
- Prompt name: `tool_usage_guide` in `prompts/tool_usage.go`

These describe when to call each tool and provide canonical usage examples.

## Server Logging

- Runtime logs are emitted via structured logging to `stderr` only.
- `LOG_LEVEL` controls verbosity and supports: `debug`, `info`, `warn`/`warning`, `error`.
- Unknown `LOG_LEVEL` values fall back to `info`.
