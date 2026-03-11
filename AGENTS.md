# Agent Guide: golang-lua-integration

## Project Overview

A Go CLI application that executes Lua scripts with Go-based tools and utilities. This implements Clean Architecture principles with domain services orchestrating Lua scripts that call Go-implemented tools.

**Key Stack**: Go 1.26, gopher-lua, urfave/cli v3, testify

---

## Essential Commands

### Build
```bash
make build
# Output: ./build/bin/golang-lua-integration
```

### Test
```bash
make test
# Runs: go test ./... --race -coverpkg=./... --coverprofile=coverage.out
# Coverage file: coverage.out
```

### Run
```bash
make run
# Creates symlink ../data -> current directory, then runs app
```

### Lint
```bash
make lint          # Run golangci-lint
make lint-fix      # Auto-fix linting issues
```

---

## Code Architecture

### Directory Structure
```
internal/
├── di/              # Dependency injection & CLI setup
├── domain/
│   ├── services/    # Business logic (ListOrderService, ListCapsService)
│   ├── model/       # Data models
│   └── tools/       # Go tools callable from Lua
├── adapter/
│   ├── cli/         # CLI command definitions
│   └── lua/         # Lua executor & file loader
├── scripts/         # Lua script files (.lua)
└── util/            # Utility functions

cmd/app/            # Application entry point (main.go)
```

### Clean Architecture Layers
1. **Domain**: `internal/domain/services/` + `internal/domain/model/` - Pure business logic
2. **Adapter**: `internal/adapter/` - CLI & Lua execution adapters
3. **DI Container**: `internal/di/di.go` - Wires up dependencies and CLI commands

---

## Key Patterns

### Service Pattern
Services in `internal/domain/services/` follow this pattern:
```go
type ServiceNameService struct {
    // fields
}

func NewServiceNameService(...) *ServiceNameService { ... }

func (s *ServiceNameService) Execute() ReturnType { ... }
```

Services orchestrate Lua scripts with Go tools. Examples:
- **ListOrderService**: Lists `/tmp` and orders alphabetically
- **ListCapsService**: Lists `/tmp` and converts to uppercase

### CLI Command Pattern
Commands defined in `internal/di/di.go`:
```go
{
    Name:        "command-name",
    Description: "...",
    Action: func(ctx context.Context, cmd *cli.Command) error {
        // Call service.Execute()
        return nil
    },
}
```

### Lua Integration
- Lua scripts located in `internal/scripts/`
- Lua scripts call Go-implemented tools (from `internal/domain/tools/`)
- Scripts and tools are executed via gopher-lua library

---

## Naming Conventions

- **Packages**: lowercase, derived from directory/purpose (e.g., `services`, `util`)
- **Exported types**: PascalCase (e.g., `ListOrderService`)
- **Functions/methods**: camelCase (e.g., `Execute()`, `NewListOrderService()`)
- **Variables**: snake_case (e.g., `list_items`)
- **Tests**: `*_test.go` suffix, test names start with `Test`
- **Comments**: `//` style, explain **why** not **what**

---

## Testing

### Running Tests
```bash
make test                                    # All tests with coverage
go test ./internal/domain/services/...       # Specific package
go test -v -run TestListOrder ./...          # Specific test
```

### Test Structure
- Test dependencies: testify (assert, require), faker (test data)
- Coverage tracked in `coverage.out`
- Use table-driven tests for multiple scenarios

### Known Patterns
1. Services use `Execute()` method (no error return, uses `panic()`)
2. Tests in same package as source code
3. Use `testify/assert` for assertions and `testify/require` for prerequisites

---

## Development Workflow

### Adding a New Command
1. Implement service in `internal/domain/services/` with `Execute()` method
2. Add command definition in `internal/di/di.go`
3. Implement or reuse Go tools in `internal/domain/tools/`
4. Create Lua script in `internal/scripts/` that uses those tools
5. Add tests following existing patterns
6. Run `make lint` and `make test` before committing

### Common Tasks
- **Add new service**: Follow `ServiceNameService` pattern with `Execute()`
- **Fix linting**: Use `make lint-fix` for auto-fixes
- **Check coverage**: Review `coverage.out` after `make test`
- **Update CLI commands**: Edit `internal/di/di.go` commands slice

---

## Dependencies

Core:
- `github.com/urfave/cli/v3` - CLI framework
- `github.com/yuin/gopher-lua` - Lua execution engine

Testing:
- `github.com/stretchr/testify` - Assertions & test utilities
- `github.com/jaswdr/faker/v2` - Test data generation

---

## Notes for Agents

- Services don't return errors; they `panic()` on failure
- Lua scripts are the orchestrators; Go code provides tools
- Keep domain logic pure; adapters handle I/O
- Always run full test suite before concluding work
- Use `make lint` to catch common issues automatically
