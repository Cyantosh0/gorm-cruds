package config

import (
	"errors"
	"fmt"

	"github.com/Cyantosh0/gorm-crud/lib"
	"github.com/Cyantosh0/gorm-crud/models"
	"gorm.io/gorm"
)

type Seed struct {
	db  Database
	env *lib.Env
}

func NewSeed(db Database, env *lib.Env) Seed {
	return Seed{
		db:  db,
		env: env,
	}
}

func (s Seed) SeedAdminUser() {
	err := s.db.First(&models.User{}, "email = ?", s.env.AdminEmail).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		s.db.Create(
			&models.User{
				Name:     s.env.AdminName,
				Email:    s.env.AdminEmail,
				Password: s.env.AdminPassword,
			})
		fmt.Println("Admin User Created")
	} else if err == nil {
		fmt.Println("Admin User Already Exists")
	}
}
