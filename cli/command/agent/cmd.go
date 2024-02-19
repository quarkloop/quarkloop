package agent

import (
	"github.com/spf13/cobra"
)

func NewAgentCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "agent",
		Short: "Manage agents",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		NewAgentCreateCommand(),
		NewAgentRemoveCommand(),
		NewAgentListCommand(),
	)
	return cmd
}
