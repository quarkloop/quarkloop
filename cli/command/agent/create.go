package agent

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/quarkloop/quarkloop/cli/client"
	v1 "github.com/quarkloop/quarkloop/pkg/grpc/v1"
	grpc "github.com/quarkloop/quarkloop/pkg/grpc/v1/agent"
)

func NewAgentCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a agent in cluster",
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace, _ := cmd.Flags().GetString("namespace")
			agent, _ := cmd.Flags().GetString("name")

			return runAgentCreate(cmd.Context(), cmd, namespace, agent)
		},
	}

	flags := cmd.Flags()
	flags.String("namespace", "", "agent namespace")
	flags.String("name", "", "agent name")

	cmd.MarkFlagRequired("namespace")
	cmd.MarkFlagRequired("name")

	return cmd
}

func runAgentCreate(ctx context.Context, cmd *cobra.Command, namespace, agent string) error {
	createCmd := &grpc.CreateAgentCommand{
		Metadata: &v1.Metadata{
			Namespace: namespace,
			Name:      agent,
		},
		CreatedBy: "cli",
	}
	_, err := client.NewClient().GetAgent().CreateAgent(context.Background(), createCmd)
	if err != nil {
		cmd.Printf("[error] unable to create %s\n", err.Error())
		return err
	}

	cmd.Printf("Agent created: %s\n", agent)
	return nil
}
