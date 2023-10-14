package db

import (
	"os"

	"github.com/anti-duhring/go-grpc/internal/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initialize() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&schema.User{})
	db.AutoMigrate(&schema.Wallet{})

	return db, nil
}
