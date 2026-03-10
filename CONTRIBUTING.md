# Contributing

Thanks for contributing to `zsv-mcp`.

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
- If you change tool inputs/outputs, update [TOOL_REFERENCE.md](docs/TOOL_REFERENCE.md) and [README.md`](./README.md) in the same PR.

