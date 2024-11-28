package utils

import (
	"log"
	"os"
	"workout-tracker/models"

	_ "github.com/lib/pq"

	"github.com/jinzhu/gorm"
)

type TrackerDB struct {
	DB *gorm.DB
}

func InitDB() (*TrackerDB, error) {
	dsn := os.Getenv("DB_DSN") // Example: "host=localhost user=postgres password=1234 dbname=workout_tracker sslmode=disable"
	if dsn == "" {
		log.Println("DB_DSN not set. Using default connection string.")
		dsn = "user=samrat dbname=workout_tracker sslmode=disable"
	}

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Exercise{},
		&models.WorkoutPlan{},
	).Error
	if err != nil {
		return nil, err
	}

	log.Println("Database connected and migrated successfully.")

	return &TrackerDB{DB: db}, nil
}

func (db *TrackerDB) Close() {
	err := db.DB.Close()
	if err != nil {
		log.Printf("Error closing the database connection: %s", err.Error())
	} else {
		log.Println("Database connection closed successfully.")
	}
}
