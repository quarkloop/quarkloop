package command

import (
	"github.com/spf13/cobra"

	"github.com/quarkloop/quarkloop/cli/command/agent"
	"github.com/quarkloop/quarkloop/cli/command/cluster"
	"github.com/quarkloop/quarkloop/cli/command/database"
	"github.com/quarkloop/quarkloop/cli/command/service"
	"github.com/quarkloop/quarkloop/cli/command/system"
)

func RegisterCommands(rootCmd *cobra.Command) {
	cmdList := []*cobra.Command{
		cluster.NewNamespaceCommand(),
		agent.NewAgentCommand(),
		service.NewServiceCommand(),
		database.NewDatabaseCommand(),
		system.NewVersionCommand(),
	}

	rootCmd.AddCommand(cmdList...)
}
