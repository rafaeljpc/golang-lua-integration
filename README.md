# golang-lua-integration

A Go CLI application that executes Lua scripts with Go-based tools. Scripts written in Lua can call Go functions to perform system tasks, combining Lua's flexibility with Go's performance and type safety.

**Technology Stack**: Go 1.26, gopher-lua, urfave/cli v3

---

## Overview

This system executes Lua scripts within Go to handle customizable tasks. The architecture separates concerns:
- **Go** implements tools and orchestration logic (domain services)
- **Lua** scripts call those tools with dynamic behavior
- **Clean Architecture** keeps business logic independent from frameworks

### Key Use Cases

#### 1. List `/tmp` and Order Alphabetically (`list-order`)

```bash
golang-lua-integration list-order
```

Workflow:
1. Lua script calls Go tool to read `/tmp` directory
2. Lua script calls Go tool to sort filenames alphabetically
3. Results returned to CLI

Implementation: `ListOrderService` + `list-order.lua`

#### 2. List `/tmp` and Convert to Uppercase (`list-caps`)

```bash
golang-lua-integration list-caps
```

Workflow:
1. Lua script calls Go tool to read `/tmp` directory
2. Lua script calls Go tool to convert all filenames to uppercase
3. Results saved to `/tmp`

Implementation: `ListCapsService` + `list-caps.lua`

---

## Project Structure

```
golang-lua-integration/
├── cmd/app/
│   └── main.go                    # Application entry point
├── internal/
│   ├── di/
│   │   └── di.go                  # Dependency injection & CLI setup
│   ├── domain/
│   │   ├── services/              # Business logic (use cases)
│   │   │   ├── listorder.go        # List & sort service
│   │   │   └── listcaps.go         # List & uppercase service
│   │   ├── model/                 # Data models
│   │   └── tools/                 # Go functions callable from Lua
│   ├── adapter/
│   │   ├── cli/                   # CLI command handlers
│   │   └── lua/                   # Lua script executor & loader
│   ├── scripts/                   # Lua script files
│   │   ├── list-order.lua
│   │   └── list-caps.lua
│   └── util/                      # Utility functions
├── Makefile                       # Build, test, lint commands
├── go.mod                         # Go module definition
├── go.sum                         # Dependency checksums
└── .golangci.yml                  # Linting configuration
```

---

## Architecture

### Clean Architecture Layers

```
┌─────────────────────────────────────┐
│  CLI (urfave/cli)                   │
│  internal/adapter/cli/              │
└────────────┬────────────────────────┘
             │
┌────────────▼────────────────────────┐
│  Domain Services (Use Cases)        │
│  internal/domain/services/          │
│  - ListOrderService.Execute()       │
│  - ListCapsService.Execute()        │
└────────────┬────────────────────────┘
             │
┌────────────▼────────────────────────┐
│  Lua Scripts ◄──► Go Tools          │
│  internal/scripts/       Domain     │
│  - list-order.lua        - model/   │
│  - list-caps.lua         - tools/   │
└─────────────────────────────────────┘
```

### Key Components

**Services** (`internal/domain/services/`)
- `ListOrderService`: Reads directory and sorts files alphabetically
- `ListCapsService`: Reads directory and converts names to uppercase
- All services implement `Execute()` method pattern

**Tools** (`internal/domain/tools/`)
- Go functions exposed to Lua scripts
- Handle filesystem operations, string transformations, etc.

**Lua Scripts** (`internal/scripts/`)
- Orchestrate tool calls from Go
- Provide dynamic behavior and scripting capabilities
- Executed via gopher-lua library

**Adapters** (`internal/adapter/`)
- `cli/`: Maps CLI commands to service calls
- `lua/`: Loads and executes Lua scripts with Go bindings

**DI Container** (`internal/di/di.go`)
- Sets up CLI commands
- Wires dependencies for services
- Routes CLI commands to service actions

---

## Development

### Build
```bash
make build
# Creates: ./build/bin/golang-lua-integration
```

### Test
```bash
make test
# Runs tests with race detection and coverage
# Output: coverage.out
```

### Run
```bash
make run
# Sets up symlink and executes app
```

### Lint
```bash
make lint        # Check code style
make lint-fix    # Auto-fix linting issues
```

---

## Design Patterns

### Service Pattern
Each service represents a use case:

```go
type ServiceNameService struct {
    // configuration fields
}

func (s *ServiceNameService) Execute() ReturnType {
    // Business logic orchestrating Lua scripts and Go tools
    return result
}
```

### CLI Command Pattern
Commands defined in `internal/di/di.go`:

```go
{
    Name:        "command-name",
    Description: "What this command does",
    Action: func(ctx context.Context, cmd *cli.Command) error {
        service := NewServiceNameService()
        result := service.Execute()
        fmt.Println(result)
        return nil
    },
}
```

---

## Adding New Features

### Add a New Command

1. **Create service** in `internal/domain/services/`:
   ```go
   type MyService struct {}
   
   func (s *MyService) Execute() []string {
       // Logic here
   }
   ```

2. **Add CLI command** in `internal/di/di.go`:
   ```go
   {
       Name: "my-command",
       Action: func(ctx context.Context, cmd *cli.Command) error {
           service := &MyService{}
           result := service.Execute()
           // Output result
           return nil
       },
   }
   ```

3. **Create Lua script** in `internal/scripts/my-command.lua`:
   ```lua
   -- Call Go tools from Lua
   files = goListDir("/tmp")
   processed = goProcessFiles(files)
   return processed
   ```

4. **Implement Go tools** in `internal/domain/tools/` if needed

5. **Write tests** for the service with `*_test.go` suffix

6. **Run tests**: `make test` and `make lint`

---

## Dependencies

### Runtime
- `github.com/urfave/cli/v3` - CLI framework
- `github.com/yuin/gopher-lua` - Lua execution engine

### Development & Testing
- `github.com/stretchr/testify` - Assertions and test utilities
- `github.com/jaswdr/faker/v2` - Test data generation

---

## Notes

- Services use `Execute()` pattern with no error return (errors trigger `panic()`)
- Lua scripts are the orchestrators; Go provides the tools
- Domain logic stays pure; adapters handle I/O and routing
- Always run `make test` and `make lint` before committing
- Coverage reports available in `coverage.out` after tests

