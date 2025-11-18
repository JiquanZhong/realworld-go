package db

import (
	"log"

	"github.com/JiquanZhong/realworld-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open(("app.db")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})

	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	log.Println("database connection established")

	// Migrate the schema
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	log.Println("database migrated successfully")
}

func GetDB() *gorm.DB {
	return DB
}
