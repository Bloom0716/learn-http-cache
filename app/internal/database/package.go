package database

import (
	"github.com/Bloom0716/learn-http-cache/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository() (Repository, error) {
	dsn := "host=posgres user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return Repository{}, err
	}

	db.AutoMigrate(&model.User{})
	return Repository{DB: db}, nil
}
