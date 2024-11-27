package utils

import (
	"log"
	"workout-tracker/models"

	_ "github.com/lib/pq"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	// dsn := os.Getenv("DB_DSN") // Example: "host=localhost user=postgres password=1234 dbname=workout_tracker sslmode=disable"
	DB, err = gorm.Open("postgres", "user=samrat dbname=workout_tracker sslmode=disable")
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Run migrations
	DB.AutoMigrate(&models.User{}, &models.Workout{})
	log.Println("Database connected and migrated")
}
