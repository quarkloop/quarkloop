package command

import (
	"fmt"
	"os"
)

func Execute() {
	rootCmd := NewRootCommand()
	RegisterCommands(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
