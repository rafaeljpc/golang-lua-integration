// Package di provides dependency injection and CLI setup.
package di

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

// DependencyInjector manages application dependencies and CLI setup.
type DependencyInjector struct {
	cmd *cli.Command
}

// Init initializes the DependencyInjector with CLI commands.
func Init(_ context.Context) *DependencyInjector {
	cmd := &cli.Command{
		Name:    "golang-lua-integration",
		Version: "0.0.1",
		Commands: []*cli.Command{
			{
				Name:        "list-order",
				Description: "List /tmp and order alphabetically",
				Action: func(_ context.Context, _ *cli.Command) error {
					return nil
				},
			},
			{
				Name:        "list-caps",
				Description: "List /tmp and convert to uppercase",
				Action: func(_ context.Context, _ *cli.Command) error {
					return nil
				},
			},
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
