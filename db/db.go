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

	// 创建默认管理员账号
	createDefaultAdmin()
}

// createDefaultAdmin 创建默认管理员账号（如果不存在）
func createDefaultAdmin() {
	var count int64
	DB.Model(&models.User{}).Where("role = ?", models.RoleAdmin).Count(&count)

	// 如果没有管理员，创建默认管理员
	if count == 0 {
		admin := models.User{
			Name:  "admin",
			Email: "admin@diit.cn",
			Age:   30,
			Role:  models.RoleAdmin,
		}

		// 设置默认密码: admin123
		if err := admin.SetPassword("Diit@123"); err != nil {
			log.Printf("Warning: failed to create default admin: %v", err)
			return
		}

		if err := DB.Create(&admin).Error; err != nil {
			log.Printf("Warning: failed to create default admin: %v", err)
			return
		}

		log.Println("Default admin account created:")
		log.Println("  Email: admin@diit.cn")
		log.Println("  Password: Diit@123")
		log.Println("  Please change the password after first login!")
	}
}

func GetDB() *gorm.DB {
	return DB
}
