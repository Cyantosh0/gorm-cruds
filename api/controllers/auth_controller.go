package controllers

import (
	"github.com/Cyantosh0/gorm-crud/api/repositories"
)

type AuthController struct {
	userRepository repositories.UserRepository
}

func NewAuthController(userRepository repositories.UserRepository) AuthController {
	return AuthController{
		userRepository: userRepository,
	}
}
