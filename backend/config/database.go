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
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
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
		{ID: 1, Username: "admin", Password: "admin123", Role: "admin"},
		{ID: 2, Username: "kerani", Password: "kerani123", Role: "kerani"},
	}
	for _, u := range users {
		db.FirstOrCreate(&models.User{}, models.User{ID: u.ID, Username: u.Username, Password: u.Password, Role: u.Role})
	}
	log.Println("Users seeded!")

	items := []models.Item{
		// Laptop
		{ID: 1, Code: "LP-001", Name: "Laptop ASUS ROG Zephyrus", Price: 25000000},
		{ID: 2, Code: "LP-002", Name: "Laptop MacBook Air M2", Price: 17000000},
		{ID: 3, Code: "LP-003", Name: "Laptop Lenovo Legion 5", Price: 19000000},
		{ID: 4, Code: "LP-004", Name: "Laptop Dell XPS 13", Price: 22000000},
		{ID: 5, Code: "LP-005", Name: "Laptop Acer Swift Go", Price: 12000000},

		// Mouse
		{ID: 6, Code: "MS-001", Name: "Mouse Logitech G Pro X", Price: 1500000},
		{ID: 7, Code: "MS-002", Name: "Mouse Razer DeathAdder V3", Price: 1200000},
		{ID: 8, Code: "MS-003", Name: "Mouse SteelSeries Rival 3", Price: 600000},
		{ID: 9, Code: "MS-004", Name: "Mouse Zowie EC2-C", Price: 1100000},
		{ID: 10, Code: "MS-005", Name: "Mouse Wireless Office M170", Price: 150000},

		// Keyboard
		{ID: 11, Code: "KB-001", Name: "Keyboard Mechanical Keychron K2", Price: 1400000},
		{ID: 12, Code: "KB-002", Name: "Keyboard Razer BlackWidow", Price: 2100000},
		{ID: 13, Code: "KB-003", Name: "Keyboard VortexSeries VX5", Price: 450000},
		{ID: 14, Code: "KB-004", Name: "Keyboard Ducky One 3", Price: 1800000},

		// Monitor & Aksesoris
		{ID: 15, Code: "MN-001", Name: "Monitor LG UltraGear 24\"", Price: 2500000},
		{ID: 16, Code: "MN-002", Name: "Monitor Samsung Odyssey G5", Price: 4500000},
		{ID: 17, Code: "HS-001", Name: "Headset HyperX Cloud II", Price: 1300000},
		{ID: 18, Code: "HS-002", Name: "Headset Sony WH-1000XM5", Price: 5000000},
		{ID: 19, Code: "CP-001", Name: "CPU Cooler Noctua NH-D15", Price: 1600000},
		{ID: 20, Code: "CP-002", Name: "Webcam Logitech C922", Price: 1200000},
	}
	for _, item := range items {
		db.FirstOrCreate(&models.Item{}, models.Item{ID: item.ID, Code: item.Code, Name: item.Name, Price: item.Price})
	}
	log.Println("Item seeded!")
}