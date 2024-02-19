package database

import (
	"github.com/spf13/cobra"
)

func NewDatabaseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "db",
		Short: "Manage database",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		NewDatabaseCreateCommand(),
		NewDatabaseRemoveCommand(),
		NewDatabaseListCommand(),
	)
	return cmd
}
