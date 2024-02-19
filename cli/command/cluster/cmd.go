package cluster

import (
	"github.com/spf13/cobra"
)

func NewNamespaceCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ns",
		Short: "Manage namespace",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		NewNamespaceCreateCommand(),
		NewNamespaceRemoveCommand(),
		NewNamespaceListCommand(),
	)
	return cmd
}
