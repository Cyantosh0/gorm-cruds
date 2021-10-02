package main

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.uber.org/fx"

	"github.com/Cyantosh0/gorm-crud/config"
	user "github.com/Cyantosh0/gorm-crud/user_crud"
)

func main() {
	godotenv.Load()

	fx.New(
		fx.Options(
			config.Module,
			user.Module,
			fx.Invoke(startApp),
		),
	).Run()
}

func startApp(
	lifecycle fx.Lifecycle,
	router config.Router,
	user_route user.UserRoute,
	migrations config.Migrations,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				migrations.Migrate()
				user_route.Setup()
				go router.Run()
				return nil
			},
		},
	)
}
