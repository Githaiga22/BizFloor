package db

import (
	"fmt"
	"log"

	"group7/models"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	// Azure SQL connection details
	const (
		server   = "all4one.database.windows.net"
		port     = 1433
		user     = "bravian"
		password = "@abcd0987" 
		database = "al4one"
	)

	// Create DSN (Data Source Name) for Azure SQL
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		user, password, server, port, database)

	// Open connection to Azure SQL using GORM
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to connect to Azure SQL Database: %v", err)
	}

	// Auto Migrate the schema
	err = db.AutoMigrate(&models.User{})
	// Uncomment if you have more models to migrate
	// err = db.AutoMigrate(&models.User{}, &models.Business{}, &models.Service{}, &models.Booking{}, &models.Payment{}, &models.SavedPlace{}, &models.LoyalClient{})
	if err != nil {
		return nil, fmt.Errorf("❌ Failed to migrate database: %v", err)
	}

	log.Println("✅ Database connected and migrated successfully")
	return db, nil
}
