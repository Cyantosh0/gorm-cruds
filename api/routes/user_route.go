package routes

import (
	"github.com/Cyantosh0/gorm-crud/api/controllers"
	"github.com/Cyantosh0/gorm-crud/config"
)

type UserRoute struct {
	handler        config.Router
	userController controllers.UserController
}

func NewUserRoute(
	handler config.Router,
	userController controllers.UserController,
) UserRoute {
	return UserRoute{
		handler:        handler,
		userController: userController,
	}
}

func (u UserRoute) Setup() {
	user := u.handler.Group("/")

	user.GET("/users", u.userController.GellAllUsers)
	user.GET("/user/:id", u.userController.RetrieveUser)
	user.POST("/user", u.userController.CreateUser)
	user.PUT("/user/:id", u.userController.UpdateUser)
	user.DELETE("/user/:id", u.userController.DeleteUser)

	user.PUT("/user-basic-information/:id", u.userController.UpdateBasicInformation)
}
