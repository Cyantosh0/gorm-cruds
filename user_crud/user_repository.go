package user

import (
	"time"

	"github.com/Cyantosh0/gorm-crud/config"
	"github.com/Cyantosh0/gorm-crud/models"
)

type UserRepository struct {
	config.Database
}

func NewUserRepository(db config.Database) UserRepository {
	return UserRepository{db}
}

func (ur UserRepository) UpdateBasicInformation(user *models.User) error {
	return ur.DB.Model(user).Updates(map[string]interface{}{
		"address":    user.Address,
		"age":        user.Age,
		"updated_at": time.Now(),
	}).Error
}
