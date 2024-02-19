package database

import (
	"context"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/quarkloop/quarkloop/cli/client"
	v1 "github.com/quarkloop/quarkloop/pkg/grpc/v1"
	grpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/table"
)

func NewDatabaseRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm",
		Short: "Remove a table from agent",
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace, _ := cmd.Flags().GetString("namespace")
			agent, _ := cmd.Flags().GetString("agent")
			table, _ := cmd.Flags().GetString("table")

			return runDatabaseRemove(cmd.Context(), cmd, namespace, agent, table)
		},
	}

	flags := cmd.Flags()
	flags.String("namespace", "ns", "")
	flags.String("agent", "", "")
	flags.String("table", "", "")

	cmd.MarkFlagRequired("namespace")
	cmd.MarkFlagRequired("agent")
	cmd.MarkFlagRequired("table")

	return cmd
}

func runDatabaseRemove(ctx context.Context, cmd *cobra.Command, namespace, agent, table string) error {
	tableId, err := strconv.ParseInt(table, 10, 32)
	if err != nil {
		cmd.Printf("Error in arg: %+v\n", err)
		return err
	}

	deleteCmd := &grpc.DeleteTableByIdCommand{
		Metadata: &v1.Metadata{
			Namespace: namespace,
			Name:      agent,
		},
		TableId: int32(tableId),
	}
	_, err = client.NewClient().GetTable().DeleteTableById(context.Background(), deleteCmd)
	if err != nil {
		cmd.Printf("Error in deleting table: %+v\n", err)
		return err
	}

	cmd.Printf("Table deleted: %d\n", tableId)
	return nil
}
