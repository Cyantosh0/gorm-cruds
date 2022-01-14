package fx_provider

import (
	"github.com/Cyantosh0/gorm-crud/cmd/flags"
	"github.com/Cyantosh0/gorm-crud/config"
	"go.uber.org/fx"
)

func Module() fx.Option {
	var module = fx.Options()

	if flags.Database {
		module = fx.Provide(config.NewDatabase)
	}

	if flags.Migration {
		module = fx.Options(module, fx.Provide(config.NewMigrations))
	}

	if flags.Seeder {
		module = fx.Options(module, fx.Provide(config.NewSeed))
	}
	return module
}
