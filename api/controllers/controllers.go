package controllers

import (
	"net/http"

	"github.com/Cyantosh0/gorm-crud/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewAuthController),
	fx.Provide(NewUserController),
)

func (ac AuthController) Login(c *gin.Context) {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	err := ac.userRepository.First(&user, "email = ?", input.Email).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user with provided email does not exist",
		})
		return
	}

	if user.Password != input.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password not correct",
		})
		return
	}

	c.JSON(200, gin.H{"msg": "success"})
}
