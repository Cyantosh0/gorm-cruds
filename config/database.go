package config

import (
	"fmt"

	"github.com/Cyantosh0/gorm-crud/lib"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
	dsn string
}

func NewDatabase(env *lib.Env) Database {
	fmt.Println("CONSOLE", env.DBName)
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", env.DBUsername, env.DBPassword, env.DBHost, env.DBPort, env.DBName)

	db, err := gorm.Open(mysql.Open(url))
	if err != nil {
		fmt.Println("Status:", err)
	}

	return Database{
		DB:  db,
		dsn: url,
	}
}
