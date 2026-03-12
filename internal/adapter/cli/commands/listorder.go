// Package commands provides CLI command implementations.
package commands

import (
	"context"

	"github.com/urfave/cli/v3"
)

// ListOrderCommand is a CLI command that lists and orders files.
type ListOrderCommand struct {
	cli.Command
}

// NewListOrderCommand creates a new ListOrderCommand instance.
func NewListOrderCommand(_ context.Context) *ListOrderCommand {
	cmd := &ListOrderCommand{
		cli.Command{
			Name:        "list-order",
			Description: "List /tmp and order alphabetically",
		},
	}

	cmd.Action = cmd.Run

	return cmd
}

// Run executes the list-order command.
func (cmd *ListOrderCommand) Run(_ context.Context, _ *cli.Command) error {
	return nil
}