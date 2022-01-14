package cli

import "github.com/spf13/cobra"

var SubCommands = []*cobra.Command{
	Version,
	AdminSeeder,
}
