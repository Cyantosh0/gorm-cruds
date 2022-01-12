package main

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.uber.org/fx"

	"github.com/Cyantosh0/gorm-crud/api/controllers"
	"github.com/Cyantosh0/gorm-crud/api/repositories"
	"github.com/Cyantosh0/gorm-crud/api/routes"
	"github.com/Cyantosh0/gorm-crud/config"
)

func main() {
	godotenv.Load()

	fx.New(
		fx.Options(
			config.Module,
			routes.Module,
			controllers.Module,
			repositories.Module,
			fx.Invoke(startApp),
		),
	).Run()
}

func startApp(
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
				migrations.Migrate()
				seed.Run()
				routes.Setup()
				go router.Run()
				return nil
			},
		},
	)
}
