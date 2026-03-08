# Agent Guide: golang-lua-integration

## Project Status

This repository appears to be a new project with no source code yet. It contains only `go.mod` and `.gitignore`.

**Action required**: Add source code files before running this command to generate meaningful documentation.

---

## Essential Commands

### Build
```bash
go build -o golang-lua-integration .
```

### Test
```bash
go test ./...
```

### Run
```bash
./golang-lua-integration
```

### Lint/Typecheck
```bash
go vet ./...
# or if using a linter:
golangci-lint run
```

---

## Code Organization

- **Root directory**: `/home/rafaeljpc/repo/golang-lua-integration`
- **Source files**: Typically in the root (Go packages)
- **Configuration**: `go.mod`, `.gitignore`

---

## Naming Conventions & Style Patterns

### Go conventions observed:
- Package names are typically lowercase, derived from file name
- Exported identifiers start with uppercase letter
- Functions and methods follow `camelCase` convention
- Variables use `snake_case` for clarity
- Comments use `//` style (single-line)

---

## Testing Approach

### Patterns observed:
- Tests are run via `go test ./...`
- Test files typically end with `_test.go` suffix
- Tests can be run on specific packages or files using package name as argument to `go test`

---

## Important Gotchas

### Project-specific context:
- This is a **new project** - no source code exists yet
- The repository is named `golang-lua-integration`, suggesting the goal is integrating Go and Lua (likely via CGO or similar)
- No existing patterns to follow beyond standard Go conventions

---

## Next Steps for Agents

1. Add source code files implementing the golang-lua integration
2. Run this command again to generate updated documentation with actual commands, patterns, and conventions
