package lib

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Env struct {
	DBHost             string `mapstructure:"DB_HOST"`
	DBPort             string `mapstructure:"DB_PORT"`
	DBName             string `mapstructure:"DB_NAME"`
	DBUsername         string `mapstructure:"DB_USER"`
	DBPassword         string `mapstructure:"DB_PASSWORD"`
	MaxMultipartMemory int64  `mapstructure:"MAX_MULTIPART_MEMORY"`
	JWTSecret          string `mapstructure:"JWT_SECRET"`
	AdminName          string `mapstructure:"ADMIN_NAME"`
	AdminEmail         string `mapstructure:"ADMIN_EMAIL"`
	AdminPassword      string `mapstructure:"ADMIN_PASSWORD"`
}

var globalEnv = Env{
	MaxMultipartMemory: 10 << 20, // 10 MB
}

func GetEnv() Env {
	return globalEnv
}

func NewEnv() *Env {
	godotenv.Load()

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic("cannot read cofiguration")
	}

	err = viper.Unmarshal(&globalEnv)
	if err != nil {
		panic(fmt.Sprintf("environment cant be loaded: ", err))
	}

	return &globalEnv
}
