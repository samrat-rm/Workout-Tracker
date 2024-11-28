package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"workout-tracker/models"
	"workout-tracker/services"
)

func CreateUser(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var user models.User
		json.NewDecoder(req.Body).Decode(&user)
		err := userService.CreateUser(&user)

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Failed to create user",
			})
			log.Printf("failed to create user : %s", err.Error())
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User created successfully",
			"id":      fmt.Sprintf("%d", user.ID),
		})
		log.Printf("user created successfully with ID %d", user.ID)
	}
}

func GetUser(w http.ResponseWriter, req *http.Request) {
	return
}
func UpdateUser(w http.ResponseWriter, req *http.Request) {
	return
}
func DeleteUser(w http.ResponseWriter, req *http.Request) {

}
func AddWorkoutToUser(w http.ResponseWriter, req *http.Request) {
	return
}
func GetUserWorkouts(w http.ResponseWriter, req *http.Request) {
	return
}
func RemoveWorkoutFromUser(w http.ResponseWriter, req *http.Request) {
	return
}
