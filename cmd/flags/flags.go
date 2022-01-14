package flags

import "github.com/spf13/cobra"

var Database bool
var Migration bool
var Seeder bool

func Setup(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&Database, "database", "d", false, "run database")
	cmd.PersistentFlags().BoolVarP(&Migration, "migration", "m", false, "run migration")
	cmd.PersistentFlags().BoolVarP(&Seeder, "seed", "s", false, "run seeder")
}
