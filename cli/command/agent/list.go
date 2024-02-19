package agent

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/quarkloop/quarkloop/cli/client"
	"github.com/quarkloop/quarkloop/cli/console"
	"github.com/quarkloop/quarkloop/pkg/grpc/v1/cluster"
)

func NewAgentListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List namespace agents",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace, _ := cmd.Flags().GetString("namespace")
			return runAgentList(cmd.Context(), cmd, namespace)
		},
	}

	flags := cmd.Flags()
	flags.String("namespace", "", "agent namespace")

	return cmd
}

type IterableAgentList struct {
	agentList        []*cluster.NamespaceAgent
	includeNamespace bool
}

func (it IterableAgentList) Iter() chan []any {
	c := make(chan []any)
	go func() {
		for _, p := range it.agentList {
			if it.includeNamespace {
				c <- []any{p.Name, p.Namespace, p.Database, p.CreatedAt.AsTime().Local()}
			} else {
				c <- []any{p.Name, p.Database, p.CreatedAt.AsTime().Local()}
			}
		}
		close(c)
	}()
	return c
}

func runAgentList(ctx context.Context, cmd *cobra.Command, namespace string) error {
	query := &cluster.GetNamespaceAgentListQuery{Namespace: namespace}
	resp, err := client.NewClient().GetCluster().GetNamespaceAgentList(context.Background(), query)
	if err != nil {
		cmd.Printf("[error] unable to get list %s\n", err.Error())
		return err
	}

	if len(namespace) == 0 {
		var cols = []any{"NAME", "NAMESPACE", "DATABASE", "CREATED"}
		printConsole(resp, cols, true)
	} else {
		var cols = []any{"NAME", "DATABASE", "CREATED"}
		printConsole(resp, cols, false)
	}

	return nil
}

func printConsole(resp *cluster.GetNamespaceAgentListResponse, cols []any, includeNamespace bool) {
	var rows IterableAgentList = IterableAgentList{
		agentList:        resp.AgentList,
		includeNamespace: includeNamespace,
	}

	var format string
	if includeNamespace {
		format = "%s\t%s\t%s\t%s\n"
	} else {
		format = "%s\t%s\t%s\n"
	}
	console.Print(format, cols, rows)
}
