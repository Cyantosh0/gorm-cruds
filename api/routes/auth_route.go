package routes

import (
	"github.com/Cyantosh0/gorm-crud/api/controllers"
	"github.com/Cyantosh0/gorm-crud/config"
)

type AuthRoute struct {
	handler        config.Router
	authController controllers.AuthController
}

func NewAuthRoute(
	handler config.Router,
	authController controllers.AuthController,
) AuthRoute {
	return AuthRoute{
		handler:        handler,
		authController: authController,
	}
}

func (ar AuthRoute) Setup() {
	auth := ar.handler.Group("/")

	auth.POST("/login", ar.authController.Login)
}
