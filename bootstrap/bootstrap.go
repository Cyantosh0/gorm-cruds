package bootstrap

import (
	"context"

	"github.com/joho/godotenv"
	"go.uber.org/fx"

	"github.com/Cyantosh0/gorm-crud/api/controllers"
	"github.com/Cyantosh0/gorm-crud/api/middlewares"
	"github.com/Cyantosh0/gorm-crud/api/repositories"
	"github.com/Cyantosh0/gorm-crud/api/routes"
	"github.com/Cyantosh0/gorm-crud/api/services"
	"github.com/Cyantosh0/gorm-crud/config"
)

func StartApp() {
	godotenv.Load()

	fx.New(
		fx.Options(
			config.Module,
			routes.Module,
			middlewares.Module,
			controllers.Module,
			services.Module,
			repositories.Module,
			fx.Invoke(bootstrap),
		),
	).Run()
}

func bootstrap(
	lifecycle fx.Lifecycle,
	database config.Database,
	router config.Router,
	seed config.Seed,
	routes routes.Routes,
	migrations config.Migrations,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				// migrations.Migrate()
				routes.Setup()
				go router.Run()
				return nil
			},
		},
	)
}
