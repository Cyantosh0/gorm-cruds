package fx_provider

import (
	"github.com/Cyantosh0/gorm-crud/cmd/flags"
	"github.com/Cyantosh0/gorm-crud/config"
	"github.com/Cyantosh0/gorm-crud/lib"
	"go.uber.org/fx"
)

func Module() fx.Option {
	var module = fx.Options(fx.Provide(lib.NewEnv), fx.Provide(config.NewMigrations))

	if flags.Database {
		module = fx.Options(module, fx.Provide(config.NewDatabase))
	}

	if flags.Seeder {
		module = fx.Options(module, fx.Provide(config.NewSeed))
	}
	return module
}
