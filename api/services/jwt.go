package services

import (
	"os"
	"time"

	"github.com/Cyantosh0/gorm-crud/constants"
	"github.com/Cyantosh0/gorm-crud/lib"
	"github.com/golang-jwt/jwt"
)

type JWTService struct{}

func NewJWTService() JWTService {
	return JWTService{}
}

//Creating Access Token
func (JWTService) CreateToken(userId lib.BinaryUUID, email string, isAdmin bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		constants.UID:   userId,
		constants.Email: email,
		constants.Admin: isAdmin,
		constants.Exp:   time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return tokenString, err
}

func (JWTService) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(tokenString *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
