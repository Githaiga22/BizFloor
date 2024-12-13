package db

import (
	"fmt"
	"log"

	"group7/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	// MySQL connection details
	const (
		DBUser     = "newuser"
		DBPassword = "password"
		DBHost     = "localhost"
		DBPort     = "3306"
		DBName     = "one4all"
	)

	// Create DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUser, DBPassword, DBHost, DBPort, DBName)

	// Open connection to MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Auto Migrate the schema
	err = db.AutoMigrate(&models.User{})
	// , &models.Business{}, &models.Service{},
	// 	&models.Booking{}, &models.Payment{}, &models.SavedPlace{}, &models.LoyalClient{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	log.Println("Database connected and migrated successfully")
	return db, nil
} 