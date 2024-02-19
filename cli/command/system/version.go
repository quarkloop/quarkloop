package system

import (
	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Display version of this CLI",
		Long:  `Display version of this CLI`,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("CLI Version: v1.3\n")
		},
	}

	return cmd
}
