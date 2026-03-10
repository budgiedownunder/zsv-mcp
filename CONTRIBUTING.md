# Contributing

Thanks for contributing to `zsv-mcp`.

## Prerequisites

- Go 1.23+
- `zsv` CLI available on your machine (or set `ZSV_PATH`)

## Local Development

1. Build the server:

```bash
go build -o zsv-mcp
```

2. Run tests:

```bash
go test ./...
```

3. Run the server manually:

```bash
./zsv-mcp
```

For VS Code MCP integration details, see `SETUP.md`.

## Code Guidelines

- Keep tool handlers small and focused.
- Validate inputs explicitly and return clear errors.
- Preserve context cancellation checks in handlers.
- Prefer table-driven tests for new behavior where practical.
- Keep public function and type names descriptive.
- For runtime logging, write to `stderr` only; do not write log output to `stdout` in this stdio MCP server.

## Testing Expectations

- Add or update tests with every behavior change.
- At minimum, test:
  - happy path
  - validation failures
  - output shape and key fields
- Ensure `go test ./...` passes before opening a PR.

## Pull Requests

- Keep PRs scoped to one logical change.
- Include a short description of:
  - what changed
  - why it changed
  - how it was tested
- If you change tool inputs/outputs, update `docs/TOOL_REFERENCE.md` and `README.md` in the same PR.

## Versioning and Module Path

- If you publish under a different GitHub org/user, update `go.mod` module path from:
  - `github.com/budgiedownunder/zsv-mcp`
  to your final repository path.
