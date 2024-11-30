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

func WriteSuccessResponse(w http.ResponseWriter, statusCode int, message string, userID *uint, token *string) {
	response := make(map[string]interface{})
	response["message"] = message

	if userID != nil {
		response["id"] = fmt.Sprintf("%d", *userID)
	}

	if token != nil {
		response["token"] = *token
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding success response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func WriteSuccessUserResponse(w http.ResponseWriter, statusCode int, user *models.User) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Printf("Error encoding user response: %v", err)
		http.Error(w, "Failed to encode user response", http.StatusInternalServerError)
	}
}
