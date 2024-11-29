package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"workout-tracker/models"
)

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	log.Println(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}

func WriteSuccessResponse(w http.ResponseWriter, statusCode int, message string, userID uint, token string) {
	log.Println(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if userID != 0 {
		json.NewEncoder(w).Encode(map[string]string{
			"message": message,
			"id":      fmt.Sprintf("%d", userID),
		})
	} else {
		json.NewEncoder(w).Encode(map[string]string{
			"token":   token,
			"message": message,
		})
	}
}

func WriteSuccessUserResponse(w http.ResponseWriter, statusCode int, user *models.User) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(user)
}
