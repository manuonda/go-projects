package database

import (
	"github.com/manuonda/projects/api-rest/go-api-crud-postgresql/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=postgres dbname=test_database port=5432 sslmode=disable timezone=Asia/Shanghai"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error connection database")
	}

	database.AutoMigrate(&models.Post{})
	database.AutoMigrate(&models.Comment{})
	DB = database
}
