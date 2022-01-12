package user

import (
	"net/http"

	"github.com/Cyantosh0/gorm-crud/lib"
	"github.com/Cyantosh0/gorm-crud/models"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	repository UserRepository
}

func NewUserController(repository UserRepository) UserController {
	return UserController{
		repository: repository,
	}
}

func (uc UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := uc.repository.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": user})
}

func (u UserController) UpdateUser(c *gin.Context) {
	paramID := c.Param("id")

	var user models.User
	err := u.repository.First(&user, "id = ?", lib.ParseUUID(paramID)).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := u.repository.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": user})
}

func (u UserController) RetrieveUser(c *gin.Context) {
	paramID := c.Param("id")

	var user models.User
	err := u.repository.First(&user, "id = ?", lib.ParseUUID(paramID)).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": user})
}

func (u UserController) GellAllUsers(c *gin.Context) {
	var users []models.User
	err := u.repository.Find(&users).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": users})
}

func (u UserController) DeleteUser(c *gin.Context) {
	paramID := c.Param("id")

	err := u.repository.Where("id = ?", lib.ParseUUID(paramID)).Delete(&models.User{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "user deleted"})
}

func (u UserController) UpdateBasicInformation(c *gin.Context) {
	paramID := c.Param("id")

	var user models.User
	err := u.repository.First(&user, "id = ?", lib.ParseUUID(paramID)).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = u.repository.UpdateBasicInformation(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": user})
}
