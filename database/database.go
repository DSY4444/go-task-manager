package database

import (
	"log"

	"github.com/DSY4444/go-task-manager/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=postgres.railway.internal user=postgres password=kKARAGhtNxQyxtuRSOSwomxTjdaVhxfh dbname=railway port=5432 sslmode=disable"
	// dsn := "host=localhost user=ds password='' dbname=taskdb port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database!")
	}
	DB.AutoMigrate(&models.Task{})
}
