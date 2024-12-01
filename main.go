package main

import (
	"log"
	"net/http"
	"workout-tracker/routes"
	"workout-tracker/services"
	"workout-tracker/utils"
)

func main() {
	trackerDB, err := utils.InitDB()
	if err != nil {
		log.Fatalln("DB connection failed, error message : ", err.Error())
	}
	defer trackerDB.Close()

	// Initialize services
	userService := services.NewUserService(trackerDB.DB)
	workoutService := services.NewWorkoutPlanService(trackerDB.DB)
	exerciseService := services.NewExerciseService(trackerDB.DB)

	// Create and set up the router
	appRouter := routes.NewAppRouter(userService, workoutService, exerciseService)

	// Start the server
	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", appRouter.Router)
}
