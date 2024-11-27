package main

import (
	"log"
	"net/http"
	"workout-tracker/handlers"
	"workout-tracker/utils"

	"github.com/gorilla/mux"
)

func main() {
	utils.InitDB()

	router := mux.NewRouter()
	router.HandleFunc("/signup", handlers.SignUp).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/logout", handlers.Logout).Methods("POST")

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
