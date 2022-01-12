package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Cyantosh0/gorm-crud/api/services"
	"github.com/Cyantosh0/gorm-crud/constants"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthMiddleware struct {
	jwtService services.JWTService
}

func NewAuthMiddleware(jwtService services.JWTService) AuthMiddleware {
	return AuthMiddleware{
		jwtService: jwtService,
	}
}

func (am AuthMiddleware) HandleAuth(c *gin.Context) {
	token, err := am.verifyToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if claims[constants.UID] != "" {
			c.Set(constants.UID, claims[constants.UID])
		}
	}

	c.Next()
}

func (am AuthMiddleware) HandleAdminOnly(c *gin.Context) {
	token, err := am.verifyToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if claims[constants.Admin] != "true" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "you are not authorized",
			})
			c.Abort()
			return
		}
	}

	c.Next()
}

func (am AuthMiddleware) verifyToken(c *gin.Context) (*jwt.Token, error) {
	header := c.GetHeader("Authorization")
	tokenString := strings.TrimSpace(strings.Replace(header, "Bearer", "", 1))

	if tokenString == "" {
		return nil, errors.New("token string is empty")
	}

	token, err := am.jwtService.VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	return token, nil
}
