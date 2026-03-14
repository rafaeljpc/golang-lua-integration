// Package commands provides CLI command implementations.
package commands

import (
	"context"
	"errors"
	"log"

	"github.com/urfave/cli/v3"
	"golang-lua-integration/internal/domain/services"
)

var errServiceNotInjected = errors.New("service is not injected")

// ListOrderCommand is a CLI command that lists and orders files.
type ListOrderCommand struct {
	cli.Command
	service services.ListService
}

// NewListOrderCommand creates a new ListOrderCommand instance with injected service.
func NewListOrderCommand(_ context.Context, service services.ListService) *ListOrderCommand {
	cmd := &ListOrderCommand{
		cli.Command{
			Name:        "list-order",
			Description: "List /tmp and order alphabetically",
		},
		service,
	}

	cmd.Action = cmd.Run

	return cmd
}

// Run executes the list-order command.
func (cmd *ListOrderCommand) Run(_ context.Context, _ *cli.Command) error {
	if cmd.service == nil {
		return errServiceNotInjected
	}

	result := cmd.service.Execute()
	for _, item := range result {
		log.Println(item)
	}

	return nil
}