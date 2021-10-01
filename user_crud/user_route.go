package user

import "github.com/Cyantosh0/gorm-crud/config"

type UserRoute struct {
	handler        config.Router
	userController UserController
}

func NewUserRoute(
	handler config.Router,
	userController UserController,
) UserRoute {
	return UserRoute{
		handler:        handler,
		userController: userController,
	}
}

func (u UserRoute) Setup() {
	user := u.handler.Group("/")

	user.POST("/user", u.userController.CreateUser)
	user.PUT("/user/:id", u.userController.UpdateUser)
}
