package db

import (
	"go-crud-api/pkg/common/models"
	model_cat "go-crud-api/pkg/models/cats"
	model_user "go-crud-api/pkg/models/users"

	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&model_cat.Cats{})
	db.AutoMigrate(&model_user.Users{})

	return db
}
