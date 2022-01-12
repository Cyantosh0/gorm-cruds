package routes

import (
	"github.com/Cyantosh0/gorm-crud/api/controllers"
	"github.com/Cyantosh0/gorm-crud/api/middlewares"
	"github.com/Cyantosh0/gorm-crud/config"
)

type AuthRoute struct {
	handler        config.Router
	authController controllers.AuthController
	authMiddleware middlewares.AuthMiddleware
}

func NewAuthRoute(
	handler config.Router,
	authController controllers.AuthController,
	authMiddleware middlewares.AuthMiddleware,
) AuthRoute {
	return AuthRoute{
		handler:        handler,
		authController: authController,
		authMiddleware: authMiddleware,
	}
}

func (ar AuthRoute) Setup() {
	auth := ar.handler.Group("/")

	auth.POST("/login", ar.authController.Login)

	auth.GET("/auth", ar.authMiddleware.HandleAuth, ar.authController.AuthenticatedController)
}
