package image

import (
	"github.com/spf13/cobra"
)

func NewImageCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "image",
		Short: "Manage images",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		NewImageBuildCommand(),
		NewImageCreateCommand(),
		NewImageRemoveCommand(),
		NewImageListCommand(),
	)
	return cmd
}
