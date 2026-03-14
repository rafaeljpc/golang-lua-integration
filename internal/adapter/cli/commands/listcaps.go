// Package commands provides CLI command implementations.
package commands

import (
	"context"
	"fmt"

	"golang-lua-integration/internal/domain/services"

	"github.com/urfave/cli/v3"
)

// ListCapsCommand is a CLI command that lists files and converts them to uppercase.
type ListCapsCommand struct {
	cli.Command
	service services.ListService
}

// NewListCapsCommand creates a new ListCapsCommand instance with injected service.
func NewListCapsCommand(_ context.Context, service services.ListService) *ListCapsCommand {
	cmd := &ListCapsCommand{
		cli.Command{
			Name:        "list-caps",
			Description: "List /tmp and convert to uppercase",
		},
		service,
	}

	cmd.Action = cmd.Run

	return cmd
}

// Run executes the list-caps command.
func (cmd *ListCapsCommand) Run(_ context.Context, _ *cli.Command) error {
	if cmd.service == nil {
		return errServiceNotInjected
	}

	result := cmd.service.Execute()
	for _, item := range result {
		fmt.Println(item)
	}

	return nil
}
