package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter() Router {
	r := gin.Default()

	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "📺 API Up and Running"})
	})

	return Router{
		r,
	}
}
