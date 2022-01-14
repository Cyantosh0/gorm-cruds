package cli

import (
	"errors"
	"fmt"

	"github.com/Cyantosh0/gorm-crud/cmd/flags"
	"github.com/Cyantosh0/gorm-crud/cmd/fx_provider"
	"github.com/Cyantosh0/gorm-crud/config"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var AdminSeeder = &cobra.Command{
	Use:   "seed-admin",
	Short: "command to seed admin user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("--- Running Admin Seeder ---")
		fx.New(
			fx_provider.Module(),
			fx.Invoke(runAdminSeeder),
		)
		fmt.Println("--- Closing Admin Seeder ---")
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkRequiredFlags()
	},
}

func checkRequiredFlags() error {
	if !flags.Database || !flags.Seeder {
		return errors.New("-d -m -s flags are required")
	}
	return nil
}

func runAdminSeeder(migration config.Migrations, seed config.Seed) {
	seed.SeedAdminUser()
}
