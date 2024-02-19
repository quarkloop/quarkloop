package database

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/quarkloop/quarkloop/cli/builder"
)

func NewDatabaseCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "Create a table for an agent",
		Aliases: []string{"apply"},
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace, _ := cmd.Flags().GetString("namespace")
			agent, _ := cmd.Flags().GetString("name")

			return runDatabaseCreate(cmd.Context(), cmd, filename)
		},
	}

	flags := cmd.Flags()
	flags.String("namespace", "", "agent namespace")
	flags.String("name", "", "agent name")

	cmd.MarkFlagRequired("namespace")
	cmd.MarkFlagRequired("name")

	return cmd
}

func runDatabaseCreate(ctx context.Context, cmd *cobra.Command, filename string) error {
	dispatcher := builder.NewDispatcher()
	err := dispatcher.ReadFile(filename)
	if err != nil {
		cmd.Printf("error on reading Quarkfile: %s", err.Error())
		return err
	}

	err = dispatcher.Dispatch()
	if err != nil {
		cmd.Printf("error on applying instructions: %s", err.Error())
		return err
	}

	return nil
}
