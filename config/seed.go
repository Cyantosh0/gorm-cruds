package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/Cyantosh0/gorm-crud/models"
	"gorm.io/gorm"
)

type Seed struct {
	db Database
}

func NewSeed(db Database) Seed {
	return Seed{
		db: db,
	}
}

func (s Seed) SeedAdminUser() {
	err := s.db.First(&models.User{}, "email = ?", os.Getenv("ADMIN_EMAIL")).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		s.db.Create(
			&models.User{
				Name:     os.Getenv("ADMIN_NAME"),
				Email:    os.Getenv("ADMIN_EMAIL"),
				Password: os.Getenv("ADMIN_PASSWORD"),
			})
		fmt.Println("Admin User Created")
	} else if err == nil {
		fmt.Println("Admin User Already Exists")
	}
}
