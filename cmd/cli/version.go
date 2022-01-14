package cli

import (
	"fmt"

	"github.com/Cyantosh0/gorm-crud/constants"
	"github.com/spf13/cobra"
)

var Version = &cobra.Command{
	Use:   "version",
	Short: "Show version of the app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(constants.Version)
	},
}
