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

	// Public routes
	router.HandleFunc("/signup", handlers.SignUp(userService)).Methods("POST")
	router.HandleFunc("/login", handlers.Login(userService)).Methods("POST")
	router.HandleFunc("/logout", handlers.Logout(userService)).Methods("POST")

	// Auth Middleware routes
	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.Use(middlewares.JwtMiddleware)
	userRouter.HandleFunc("/{id}", handlers.GetUser(userService)).Methods("GET")
	userRouter.HandleFunc("/{id}", handlers.DeleteUser(userService)).Methods("DELETE")
	userRouter.HandleFunc("/{id}", handlers.UpdateUser(userService)).Methods("POST")

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
