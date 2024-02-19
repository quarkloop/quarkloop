package cluster

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/quarkloop/quarkloop/cli/client"
	"github.com/quarkloop/quarkloop/pkg/grpc/v1/cluster"
)

func NewNamespaceCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a namespace in cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace, _ := cmd.Flags().GetString("name")
			return runNamespaceCreate(cmd.Context(), cmd, namespace)
		},
	}

	flags := cmd.Flags()
	flags.String("name", "", "namespace name")

	cmd.MarkFlagRequired("name")
	return cmd
}

func runNamespaceCreate(ctx context.Context, cmd *cobra.Command, namespace string) error {
	createCmd := &cluster.CreateNamespaceCommand{Namespace: namespace}
	_, err := client.NewClient().GetCluster().CreateNamespace(context.Background(), createCmd)
	if err != nil {
		cmd.Printf("[error] unable to create %s\n", err.Error())
		return err
	}

	cmd.Printf("Namespace created: %s\n", namespace)
	return nil
}
