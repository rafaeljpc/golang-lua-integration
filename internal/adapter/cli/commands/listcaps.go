package commands

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

type ListCapsCommand struct {
	cli.Command
}

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

func (cmd *ListCapsCommand) Run(ctx context.Context, command *cli.Command) error {
	fmt.Printf("list-caps called\n")

	return nil
}
