package services

import (
	"time"

	"github.com/Cyantosh0/gorm-crud/constants"
	"github.com/Cyantosh0/gorm-crud/lib"
	"github.com/golang-jwt/jwt"
)

type JWTService struct {
	env *lib.Env
}

func NewJWTService(env *lib.Env) JWTService {
	return JWTService{
		env: env,
	}
}

//Creating Access Token
func (j JWTService) CreateToken(userId lib.BinaryUUID, email string, isAdmin bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		constants.UID:   userId,
		constants.Email: email,
		constants.Admin: isAdmin,
		constants.Exp:   time.Now().Add(time.Minute * 15).Unix(),
	})

	tokenString, err := token.SignedString([]byte(j.env.JWTSecret))
	return tokenString, err
}

func (j JWTService) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(tokenString *jwt.Token) (interface{}, error) {
		return []byte(j.env.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
