package commands

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

type ListOrderCommand struct {
	cli.Command
}

func NewListOrderCommand(ctx context.Context) *ListOrderCommand {
	cmd := &ListOrderCommand{
		cli.Command{
			Name:        "list-order",
			Description: "List /tmp and order alphabetically",
		},
	}

	cmd.Action = cmd.Run

	return cmd
}

func (cmd *ListOrderCommand) Run(ctx context.Context, command *cli.Command) error {
	fmt.Printf("list-order called\n")
	
	return nil
}