package main

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.uber.org/fx"

	"github.com/Cyantosh0/gorm-crud/config"
)

func main() {
	godotenv.Load()

	fx.New(
		fx.Options(config.Module),
		fx.Invoke(startApp),
	).Run()
}

func startApp(
	lifecycle fx.Lifecycle,
	router config.Router,
	migrations config.Migrations,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				migrations.Migrate()
				_ = router.Run()
				return nil
			},
		},
	)
}
