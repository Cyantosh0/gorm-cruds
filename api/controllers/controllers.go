package controllers

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewAuthController),
	fx.Provide(NewUserController),
	fx.Provide(NewFileController),
)
