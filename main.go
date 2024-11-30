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

	authHandler := handlers.NewAuthHandler(userService)
	userHandler := handlers.NewUserHandler(userService)

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

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
