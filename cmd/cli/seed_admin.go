package cli

import (
	"fmt"

	"github.com/Cyantosh0/gorm-crud/config"
	"github.com/spf13/cobra"
)

func AdminSeeder(seed config.Seed) *cobra.Command {
	return &cobra.Command{
		Use:   "seed-admin",
		Short: "command to seed admin user",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("--- Running Admin Seeder ---")
			seed.SeedAdminUser()
			fmt.Println("--- Closing Admin Seeder ---")
		},
	}
}
