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
	db := utils.InitDB()

	router := mux.NewRouter()

	userService := services.NewUserService(db)

	// Public routes
	router.HandleFunc("/signup", handlers.SignUp(userService)).Methods("POST")
	router.HandleFunc("/login", handlers.Login(userService)).Methods("POST")
	router.HandleFunc("/logout", handlers.Logout(userService)).Methods("POST")

	// Protected routes with middleware
	protectedRoutes := router.PathPrefix("/user").Subrouter()
	protectedRoutes.Use(middlewares.JwtMiddleware)
	protectedRoutes.HandleFunc("/{id}", handlers.GetUser(userService)).Methods("GET")

	// router.HandleFunc("/user", handlers.UpdateUser).Methods("PUT")
	// router.HandleFunc("/user", handlers.DeleteUser).Methods("DELETE")
	// router.HandleFunc("/user/workout", handlers.AddWorkoutToUser).Methods("POST")
	// router.HandleFunc("/user/workout", handlers.GetUserWorkouts).Methods("GET")
	// router.HandleFunc("/user/workout", handlers.RemoveWorkoutFromUser).Methods("DELETE")

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
