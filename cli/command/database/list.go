package database

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/quarkloop/quarkloop/cli/client"
	"github.com/quarkloop/quarkloop/cli/console"
	v1 "github.com/quarkloop/quarkloop/pkg/grpc/v1"
	"github.com/quarkloop/quarkloop/pkg/grpc/v1/table"
)

func NewDatabaseListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List agent tables",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace, _ := cmd.Flags().GetString("namespace")
			agent, _ := cmd.Flags().GetString("agent")

			return runDatabaseList(cmd.Context(), cmd, namespace, agent)
		},
	}

	flags := cmd.Flags()
	flags.String("namespace", "ns", "")
	flags.String("agent", "", "")

	cmd.MarkFlagRequired("namespace")
	cmd.MarkFlagRequired("agent")

	return cmd
}

type IterableTableList []*table.Table

func (it IterableTableList) Iter() chan []any {
	c := make(chan []any)
	go func() {
		for _, p := range it {
			c <- []any{p.Name, p.Schema, p.CreatedAt.AsTime().Local()}
		}
		close(c)
	}()
	return c
}

func runDatabaseList(ctx context.Context, cmd *cobra.Command, namespace, agent string) error {
	query := &table.GetTableListQuery{
		Metadata: &v1.Metadata{
			Namespace: namespace,
			Name:      agent,
		},
	}
	resp, err := client.NewClient().GetTable().GetTableList(ctx, query)
	if err != nil {
		cmd.Printf("error in get table list: %+v\n", err)
		return err
	}

	var cols = []any{"NAME", "SCHEMA", "CREATED"}
	var rows IterableTableList = resp.Tables
	console.Print("%s\t%s\t%s\n", cols, rows)

	return nil
}
