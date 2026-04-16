package config

import (
	"fleetify-backend/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	log.Println("Database Connected!")

	err = db.AutoMigrate(
		&models.User{},
		&models.Item{},
		&models.Invoice{},
		&models.InvoiceDetail{},
	)
	if err != nil {
		log.Fatal("Failed to migate database: ", err)
	}

	log.Println("Database migrated!")

	DB = db

	seedData(db)
}

func seedData(db *gorm.DB) {
	users := []models.User{
		{Username: "Admin", Password: "@Admin123", Role: "admin"},
		{Username: "Kerani", Password: "@Kerani123", Role: "kerani"},
	}
	for _, u := range users {
		db.FirstOrCreate(&models.User{}, models.User{Username: u.Username, Password: u.Password, Role: u.Role})
	}
	log.Println("Users seeded!")

	items := []models.Item{
		{Code: "BRG-001", Name: "Mouse", Price: 750000},
		{Code: "BRG-002", Name: "Laptop", Price: 16500000},
		{Code: "BRG-003", Name: "Keyboard", Price: 850000},
		{Code: "BRG-004", Name: "Headset", Price: 1200000},
	}
	for _, item := range items {
		db.FirstOrCreate(&models.Item{}, models.Item{Code: item.Code, Name: item.Name, Price: item.Price})
	}
	log.Println("Item seeded!")
}