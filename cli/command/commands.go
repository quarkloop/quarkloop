package command

import (
	"github.com/spf13/cobra"

	"github.com/quarkloop/quarkloop/cli/command/agent"
	"github.com/quarkloop/quarkloop/cli/command/cluster"
	"github.com/quarkloop/quarkloop/cli/command/database"
	"github.com/quarkloop/quarkloop/cli/command/image"
	"github.com/quarkloop/quarkloop/cli/command/object"
	"github.com/quarkloop/quarkloop/cli/command/service"
	"github.com/quarkloop/quarkloop/cli/command/system"
)

func RegisterCommands(rootCmd *cobra.Command) {
	registerManagementCommands(rootCmd)
	registerNormalCommands(rootCmd)
}

func registerManagementCommands(rootCmd *cobra.Command) {
	cmdList := []*cobra.Command{
		cluster.NewNamespaceCommand(),
		image.NewImageCommand(),
		agent.NewAgentCommand(),
		service.NewServiceCommand(),
		database.NewDatabaseCommand(),
	}
	for _, cmd := range cmdList {
		cmd.GroupID = "management"
	}

	group := &cobra.Group{
		ID:    "management",
		Title: "Management Commands:",
	}
	rootCmd.AddGroup(group)
	rootCmd.AddCommand(cmdList...)
}

func registerNormalCommands(rootCmd *cobra.Command) {
	cmdList := []*cobra.Command{
		image.NewImageBuildCommand(),
		object.NewInitCommand(),
		object.NewGenerateCommand(),
		object.NewApplyCommand(),
		object.NewCreateCommand(),
		object.NewDeleteCommand(),
		system.NewVersionCommand(),
	}
	for _, cmd := range cmdList {
		cmd.GroupID = "normal"
	}

	group := &cobra.Group{
		ID:    "normal",
		Title: "Commands:",
	}
	rootCmd.AddGroup(group)
	rootCmd.SetHelpCommandGroupID("normal")
	rootCmd.AddCommand(cmdList...)
}
