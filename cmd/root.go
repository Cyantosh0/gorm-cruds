package cmd

import (
	"fmt"

	"github.com/Cyantosh0/gorm-crud/bootstrap"
	"github.com/spf13/cobra"
)

var (
	Version = "1.0.0-0"

	RootCmd = &cobra.Command{
		Use:     "gorm-crud",
		Short:   "root command of our application",
		Long:    `root command of our application,the main purpose is to help to setup subcommands`,
		Version: Version,
		Run: func(cmd *cobra.Command, args []string) {
			bootstrap.StartApp()
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Show version of the app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(RootCmd.Version)
		},
	}
)

func init() {
	RootCmd.AddCommand(versionCmd)
}
