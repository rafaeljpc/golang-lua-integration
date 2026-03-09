# Agent Guide: golang-lua-integration

## Project Status

This repository implements a Go application for integrating with Lua (likely via CGO or similar). The project uses standard Go conventions and has source code in `cmd/`, `internal/`, and root directories.

---

## Essential Commands

### Build
```bash
go build -o ./build/bin/golang-lua-integration ./cmd/app
# or from Makefile:
make build
```

### Test
```bash
go test ./... --race -coverpkg=./...  --coverprofile=coverage.out
# or from Makefile:
make test
```

### Run
```bash
go run ./cmd/app
# or from Makefile:
make run
```

### Lint/Typecheck
```bash
golangci-lint run
# or from Makefile:
make lint
```

---

## Code Organization

- **Root directory**: `/home/rafaeljpc/repo/golang-lua-integration`
- **Application entry point**: `cmd/app/main.go`
- **Domain layer**: `internal/domain/services/`, `internal/domain/model/`, `internal/domain/tools/`
- **Adapter layer**: `internal/adapter/cli/`, `internal/adapter/lua/`
- **Utility layer**: `internal/util/`
- **Dependency injection**: `internal/di/di.go`

---

## Naming Conventions & Style Patterns

### Go conventions observed:
- Package names are typically lowercase, derived from file name (e.g., `services`, `util`)
- Exported identifiers start with uppercase letter
- Functions and methods follow `camelCase` convention
- Variables use `snake_case` for clarity
- Comments use `//` style (single-line)
- Service types follow pattern: `ServiceName + "Service"` struct with `Execute()` method

---

## Testing Approach & Patterns

### Patterns observed:
- Tests are run via `go test ./... --race -coverpkg=./...  --coverprofile=coverage.out`
- Test files typically end with `_test.go` suffix
- Tests can be run on specific packages or files using package name as argument to `go test`
- Coverage is tracked in `coverage.out` file

### Important Gotchas:

1. **Type compatibility errors**: Some test files have type mismatches (e.g., `*bytes.Buffer` vs `*os.File`). When fixing tests, ensure types match the expected interface.

2. **Interface{} usage**: The codebase uses `interface{}` for generic types; consider using `any` or specific types where appropriate.

3. **Logging patterns**: Uses `log.Default().Printf()` and `log.Printf()` for error logging.

4. **Error handling**: Some services use `panic(err)` for errors instead of returning them via the interface.

---

## Project-Specific Context

- The repository is named `golang-lua-integration`, suggesting the goal is integrating Go and Lua (likely via CGO or similar)
- Uses `github.com/yuin/gopher-lua` as a dependency
- Domain services implement use cases: ListOrderService, ListCapsService
- Dependency injection pattern used in main entry point

---

## Next Steps for Agents

1. When adding new domain services, follow the `Execute()` method pattern
2. Ensure type compatibility when working with test files
3. Use `golangci-lint` to catch type mismatches and other issues
4. Coverage reports are generated in `coverage.out` file
