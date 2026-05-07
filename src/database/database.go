package database

import (
	"app/src/config"
	"app/src/model"
	"app/src/utils"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(dbHost, dbName string) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, config.DBUser, config.DBPassword, dbName, config.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		TranslateError:         true,
	})
	if err != nil {
		utils.Log.Errorf("Failed to connect to database: %+v", err)
	}

	sqlDB, errDB := db.DB()
	if errDB != nil {
		utils.Log.Errorf("Failed to connect to database: %+v", errDB)
	}

	// Config connection pooling
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)



	// Run Migrations
	if err := db.AutoMigrate(&model.User{}, &model.CompareHistory{}); err != nil {
		utils.Log.Errorf("Failed to auto migrate: %+v", err)
	}

	// Seed Data if empty
	seedData(db)

	return db
}

func seedData(db *gorm.DB) {
	// Seed Default Admin User
	var adminCount int64
	db.Model(&model.User{}).Where("email = ?", "admin@admin.com").Count(&adminCount)
	if adminCount == 0 {
		hashedPassword, _ := utils.HashPassword("admin")
		adminUser := model.User{
			Name:         "Admin User",
			Email:        "admin@admin.com",
			Phone:        "01700000000",
			PasswordHash: hashedPassword,
			Status:       "active",
		}
		db.Create(&adminUser)
		utils.Log.Info("Default admin user created - Email: admin@admin.com, Password: admin")
	}
}

