package cmd

import (
	"github.com/Cyantosh0/gorm-crud/bootstrap"
	"github.com/Cyantosh0/gorm-crud/cmd/cli"
	"github.com/Cyantosh0/gorm-crud/config"
	"github.com/Cyantosh0/gorm-crud/constants"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
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
	godotenv.Load()
	RootCmd.AddCommand(cli.Version)

	fx.New(
		fx.Options(
			fx.Provide(config.NewDatabase),
			fx.Provide(config.NewMigrations),
			fx.Provide(config.NewSeed),
			fx.Invoke(runDependentCommands),
		),
	)
}

func runDependentCommands(
	migration config.Migrations,
	seed config.Seed,
) {
	migration.Migrate()
	RootCmd.AddCommand(cli.AdminSeeder(seed))
}
