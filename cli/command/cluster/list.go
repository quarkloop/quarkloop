package cluster

import (
	"context"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/quarkloop/quarkloop/cli/client"
	"github.com/quarkloop/quarkloop/cli/console"
	"github.com/quarkloop/quarkloop/pkg/grpc/v1/cluster"
)

func NewNamespaceListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List cluster namespaces",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runNamespaceList(cmd.Context(), cmd)
		},
	}

	return cmd
}

type IterableNamespaceList []*cluster.Namespace

func (it IterableNamespaceList) Iter() chan []any {
	c := make(chan []any)
	go func() {
		for _, p := range it {
			c <- []any{p.Namespace, p.CreatedAt.AsTime().Local()}
		}
		close(c)
	}()
	return c
}

func runNamespaceList(ctx context.Context, cmd *cobra.Command) error {
	resp, err := client.NewClient().GetCluster().GetNamespaceList(context.Background(), &emptypb.Empty{})
	if err != nil {
		cmd.Printf("[error] unable to get list %s\n", err.Error())
		return err
	}

	var cols = []any{"NAME", "CREATED"}
	var rows IterableNamespaceList = resp.NamespaceList
	console.Print("%s\t%s\n", cols, rows)

	return nil
}
