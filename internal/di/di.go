package di

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

type DependencyInjector struct {
	cmd *cli.Command
}

func Init(ctx context.Context) *DependencyInjector {
	cmd := &cli.Command{
		Name:    "golang-lua-integration",
		Version: "0.0.1",
		Commands: []*cli.Command{
			{
				Name:        "list-order",
				Description: "List /tmp and order alphabetically",
				Action: func(ctx context.Context, command *cli.Command) error {
					fmt.Printf("list-order called\n")

					return nil
				},
			},
			{
				Name:        "list-caps",
				Description: "List /tmp and convert to uppercase",
				Action: func(ctx context.Context, command *cli.Command) error {
					fmt.Printf("list-caps called\n")

					return nil
				},
			},
		},
	}

	return &DependencyInjector{
		cmd: cmd,
	}
}

func (di *DependencyInjector) Run(ctx context.Context, args []string) {
	di.cmd.Run(context.Background(), os.Args)
}
