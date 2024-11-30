package main

import (
	"log"
	"net/http"
	"workout-tracker/handlers"
	"workout-tracker/middlewares"
	"workout-tracker/services"
	"workout-tracker/utils"

	"github.com/gorilla/mux"
)

func main() {
	trackerDB, err := utils.InitDB()
	if err != nil {
		log.Fatalln("DB connection failed, error message : ", err.Error())
	}
	defer trackerDB.Close()

	router := mux.NewRouter()

	userService := services.NewUserService(trackerDB.DB)
	workoutService := services.NewWorkoutPlanService(trackerDB.DB)

	authHandler := handlers.NewAuthHandler(userService)
	userHandler := handlers.NewUserHandler(userService)
	workoutPlanHandler := handlers.NewWokoutSericeHandler(workoutService, userService)

	// Public routes
	router.HandleFunc("/signup", authHandler.SignUp).Methods("POST")
	router.HandleFunc("/login", authHandler.Login).Methods("POST")
	router.HandleFunc("/logout", authHandler.Logout).Methods("POST")

	// Auth Middleware routes
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.Use(middlewares.JwtMiddleware)
	userRouter.HandleFunc("/{id}", userHandler.GetUser).Methods("GET")
	userRouter.HandleFunc("/{id}", userHandler.DeleteUser).Methods("DELETE")
	userRouter.HandleFunc("/{id}", userHandler.UpdateUser).Methods("POST")

	userRouter.HandleFunc("/{id}/workout_plan", workoutPlanHandler.GetAllWorkoutPlansForUser).Methods("GET") // /user/{id}/workout_plan
	userRouter.HandleFunc("/{id}/workout_plan/{wp_id}", workoutPlanHandler.GetWorkoutPlanForUser).Methods("GET")

	workoutPlanRouter := router.PathPrefix("/workout_plan").Subrouter()
	workoutPlanRouter.Use(middlewares.JwtMiddleware)
	workoutPlanRouter.HandleFunc("/{id}", workoutPlanHandler.CreateWorkoutPlanForUser).Methods("POST")

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
