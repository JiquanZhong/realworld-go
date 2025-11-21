package db

import (
	"os"

	"github.com/JiquanZhong/realworld-go/models"
	"github.com/JiquanZhong/realworld-go/utils"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	var err error

	driver := os.Getenv("DB_DRIVER")
	if driver == "" {
		driver = "sqlite"
	}

	switch driver {
	case "postgres", "pg":
		dsn := os.Getenv("DB_DSN")
		if dsn == "" {
			utils.Logger.Fatal("DB_DSN is not set")
		}
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
		})
	default:
		DB, err = gorm.Open(sqlite.Open(("app.db")), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
		})
	}

	if err != nil {
		utils.Logger.Fatal("failed to connect database: ", zap.Error(err))
	}

	utils.Logger.Info("database connection established")

	// Migrate the schema
	if err := DB.AutoMigrate(
		&models.User{},
		&models.McpService{},
		&models.McpTag{},
		&models.MCPServiceTag{},
		&models.McpRating{},
		&models.McpFavorite{},
		&models.McpInstallation{},
		&models.McpToken{},
	); err != nil {
		utils.Logger.Fatal("failed to migrate database: ", zap.Error(err))
	}

	utils.Logger.Info("database migrated successfully")

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
			utils.Logger.Warn("Warning: failed to create default admin: ", zap.Error(err))
			return
		}

		if err := DB.Create(&admin).Error; err != nil {
			utils.Logger.Warn("Warning: failed to create default admin: ", zap.Error(err))
			return
		}

		utils.Logger.Info("Default admin account created:")
		utils.Logger.Info("  Email: admin@diit.cn")
		utils.Logger.Info("  Password: Diit@123")
		utils.Logger.Info("  Please change the password after first login!")
	}
}

func GetDB() *gorm.DB {
	return DB
}
