package user

import (
	"net/http"

	"github.com/Cyantosh0/gorm-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
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
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := uc.repository.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{"data": user})
}

func (u UserController) UpdateUser(c *gin.Context) {
	paramID := c.Param("id")

	type inputType struct {
		models.User
		ShouldUpdate bool `json:"shouldUpdate" form:"shouldUpdate"`
	}
	var input inputType
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	if input.ShouldUpdate {
		err := u.repository.First(&user, "id = ?", paramID).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		_ = copier.Copy(&user, input)

		if err := u.repository.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}
	}

	c.JSON(200, gin.H{"data": user})
}

func (u UserController) RetrieveUser(c *gin.Context) {
	paramID := c.Param("id")

	var user models.User
	err := u.repository.First(&user, "id = ?", paramID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
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
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{"data": users})
}

func (u UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := u.repository.Delete(&models.User{}, id).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{"data": "user deleted"})
}

func (u UserController) UpdateBasicInformation(c *gin.Context) {
	paramID := c.Param("id")

	var user models.User
	err := u.repository.First(&user, "id = ?", paramID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
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
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{"data": user})
}
