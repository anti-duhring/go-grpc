package db

import (
	"os"

	"github.com/anti-duhring/go-grpc/internal/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Client *gorm.DB

func Initialize() error {
	dsn := os.Getenv("DATABASE_URL")

	var err error
	Client, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})

	if err != nil {
		return err
	}

	Client.AutoMigrate(&schema.User{})
	Client.AutoMigrate(&schema.Wallet{})

	return nil
}
