package cluster

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/quarkloop/quarkloop/cli/client"
	cluster "github.com/quarkloop/quarkloop/pkg/grpc/daemon/v1/cluster"
)

func NewNamespaceRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm",
		Short: "Remove a namespace in cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace, _ := cmd.Flags().GetString("name")
			return runNamespaceRemove(cmd.Context(), cmd, namespace)
		},
	}

	flags := cmd.Flags()
	flags.String("name", "", "namespace name to remove")

	cmd.MarkFlagRequired("name")
	return cmd
}

func runNamespaceRemove(ctx context.Context, cmd *cobra.Command, namespace string) error {
	deleteCmd := &cluster.DeleteNamespaceCommand{Namespace: namespace}
	_, err := client.NewClient().GetCluster().DeleteNamespace(context.Background(), deleteCmd)
	if err != nil {
		cmd.Printf("[error] unable to create %s\n", err.Error())
		return err
	}

	cmd.Printf("Namespace deleted: %s\n", namespace)
	return nil
}
