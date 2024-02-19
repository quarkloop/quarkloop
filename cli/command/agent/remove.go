package agent

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/quarkloop/quarkloop/cli/client"
	v1 "github.com/quarkloop/quarkloop/pkg/grpc/v1"
	grpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/agent"
)

func NewAgentRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm",
		Short: "Remove a agent in cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace, _ := cmd.Flags().GetString("namespace")
			agent, _ := cmd.Flags().GetString("name")

			return runAgentRemove(cmd.Context(), cmd, namespace, agent)
		},
	}

	flags := cmd.Flags()
	flags.String("namespace", "", "agent namespace")
	flags.String("name", "", "agent name to remove")

	cmd.MarkFlagRequired("namespace")
	cmd.MarkFlagRequired("name")

	return cmd
}

func runAgentRemove(ctx context.Context, cmd *cobra.Command, namespace, agent string) error {
	deleteCmd := &grpc.DeleteAgentCommand{
		Metadata: &v1.Metadata{
			Namespace: namespace,
			Name:      agent,
		},
	}
	_, err := client.NewClient().GetAgent().DeleteAgent(context.Background(), deleteCmd)
	if err != nil {
		cmd.Printf("[error] unable to create %s\n", err.Error())
		return err
	}

	cmd.Printf("Agent deleted: %s\n", agent)
	return nil
}
