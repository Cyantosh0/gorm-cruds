package cmd

import (
	"github.com/Cyantosh0/gorm-crud/bootstrap"
	"github.com/Cyantosh0/gorm-crud/cmd/cli"
	"github.com/Cyantosh0/gorm-crud/cmd/flags"
	"github.com/Cyantosh0/gorm-crud/constants"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "gorm-crud",
	Short:   "root command of our application",
	Long:    `root command of our application,the main purpose is to help to setup subcommands`,
	Version: constants.Version,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.StartApp()
	},
}

func init() {
	flags.Setup(RootCmd)

	for _, cmd := range cli.SubCommands {
		RootCmd.AddCommand(cmd)
	}
}
