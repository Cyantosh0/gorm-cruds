package user

import "github.com/Cyantosh0/gorm-crud/config"

type UserRepository struct {
	config.Database
}

func NewUserRepository(db config.Database) UserRepository {
	return UserRepository{db}
}
