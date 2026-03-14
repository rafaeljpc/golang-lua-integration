// Package di provides dependency injection and CLI setup.
package di

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
	"golang-lua-integration/internal/adapter/cli/commands"
	"golang-lua-integration/internal/domain/services"
)

// DependencyInjector manages application dependencies and CLI setup.
type DependencyInjector struct {
	cmd *cli.Command
}

// Init initializes the DependencyInjector with CLI commands.
func Init(ctx context.Context) *DependencyInjector {
	listOrderService := services.NewListOrderService("/tmp")
	listCapsService := services.NewListCapsService("/tmp")

	listOrderCmd := commands.NewListOrderCommand(ctx, listOrderService)
	listCapsCmd := commands.NewListCapsCommand(ctx, listCapsService)

	cmd := &cli.Command{
		Name:    "golang-lua-integration",
		Version: "0.0.1",
		Commands: []*cli.Command{
			&listOrderCmd.Command,
			&listCapsCmd.Command,
		},
	}

	return &DependencyInjector{
		cmd: cmd,
	}
}

// Run executes the CLI application.
func (di *DependencyInjector) Run(ctx context.Context, _ []string) {
	err := di.cmd.Run(ctx, os.Args)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
