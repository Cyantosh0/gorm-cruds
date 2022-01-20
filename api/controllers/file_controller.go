package controllers

import (
	"github.com/gin-gonic/gin"
)

type FileController struct {
}

func NewFileController() FileController {
	return FileController{}
}

func (fc FileController) UploadFile(c *gin.Context) {

	c.JSON(200, gin.H{"msg": "success"})
}
