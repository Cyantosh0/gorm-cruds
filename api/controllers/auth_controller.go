package controllers

import (
	"net/http"

	"github.com/Cyantosh0/gorm-crud/api/repositories"
	"github.com/Cyantosh0/gorm-crud/api/services"
	"github.com/Cyantosh0/gorm-crud/constants"
	"github.com/Cyantosh0/gorm-crud/lib"
	"github.com/Cyantosh0/gorm-crud/models"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userRepository repositories.UserRepository
	jwtService     services.JWTService
	env            *lib.Env
}

func NewAuthController(
	userRepository repositories.UserRepository,
	jwtService services.JWTService,
	env *lib.Env,
) AuthController {
	return AuthController{
		userRepository: userRepository,
		jwtService:     jwtService,
		env:            env,
	}
}

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

	isAdmin := false
	if user.Email == ac.env.AdminEmail {
		isAdmin = true
	}

	token, err := ac.jwtService.CreateToken(user.ID, user.Email, isAdmin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, token)
}

func (ac AuthController) AuthenticatedController(c *gin.Context) {
	userUID := c.MustGet(constants.UID)
	c.JSON(http.StatusOK, gin.H{
		constants.UID: userUID,
	})
}
