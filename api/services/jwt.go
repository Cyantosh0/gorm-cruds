package services

import (
	"os"
	"time"

	"github.com/Cyantosh0/gorm-crud/lib"
	"github.com/golang-jwt/jwt"
)

//Creating Access Token
func CreateToken(userId lib.BinaryUUID, email string, isAdmin bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"email":   email,
		"admin":   isAdmin,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return tokenString, err
}
