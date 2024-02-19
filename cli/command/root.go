package command

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "quarkloop",
		Short:             "A runtime for services",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: false, HiddenDefaultCmd: true},
	}

	return cmd
}
