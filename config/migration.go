package config

import (
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

type Migrations struct {
	db Database
}

func NewMigrations(db Database) Migrations {
	return Migrations{
		db: db,
	}
}

func (m Migrations) Migrate() {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations/",
	}

	sqlDB, err := m.db.DB.DB()
	if err != nil {
		panic(err)
	}

	_, err = migrate.Exec(sqlDB, "mysql", migrations, migrate.Up)
	if err != nil {
		fmt.Println("Error in migration: ", err)
	}
}
