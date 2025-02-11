package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/webolton/prat/internal/bootstrap"
)

func newVersionCommand(config bootstrap.Config) *cobra.Command {
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of prat",
		Long:  `Prints the version number of prat`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(config.Environment)
		},
	}
	return versionCmd
}
