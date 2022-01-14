package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/Cyantosh0/gorm-crud/cmd"
)

func main() {
	cmd.RootCmd.Execute()
}
