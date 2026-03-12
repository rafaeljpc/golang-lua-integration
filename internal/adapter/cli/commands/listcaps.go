// Package commands provides CLI command implementations.
package commands

import (
	"context"

	"github.com/urfave/cli/v3"
)

// ListCapsCommand is a CLI command that lists files and converts them to uppercase.
type ListCapsCommand struct {
	cli.Command
}

// NewListCapsCommand creates a new ListCapsCommand instance.
func NewListCapsCommand(ctx context.Context) *ListCapsCommand {
	cmd := &ListCapsCommand{
		cli.Command{
			Name:        "list-caps",
			Description: "List /tmp and convert to uppercase",
		},
	}

	cmd.Action = cmd.Run

	return cmd
}

// Run executes the list-caps command.
func (cmd *ListCapsCommand) Run(_ context.Context, _ *cli.Command) error {
	return nil
}
