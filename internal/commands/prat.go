package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "prat",
	Short: "\"Pack rat for syncing files to the cloud.\"",
	Long:  "A file syncing program that runs as a service",
	RunE: func(commands *cobra.Command, args []string) error {
		return commands.Help()
	},
}

func Execute(something string) {
	rootCmd.AddCommand(newVersionCommand(something))
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
