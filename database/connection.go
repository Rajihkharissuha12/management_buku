package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Railway menyediakan DATABASE_URL
	dsn := os.Getenv("DATABASE_URL")

	// Jika run lokal, fallback ke local DSN
	if dsn == "" {
		dsn = "host=localhost user=postgres password=rajih123 dbname=managementbuku port=5432 sslmode=disable"
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected")
}
