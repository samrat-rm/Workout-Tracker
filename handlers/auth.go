package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"workout-tracker/models"
	"workout-tracker/services"
	"workout-tracker/utils"

	"golang.org/x/crypto/bcrypt"
)

func SignUp(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var user models.User
		err := json.NewDecoder(req.Body).Decode(&user)
		if err != nil || !utils.ValidateUserCredentials(&user) {
			msg := "Invalid credentials, Failed to Sign up"
			if err != nil {
				msg += ", " + err.Error()
			}
			writeErrorResponse(w, http.StatusBadRequest, msg)
			return
		}

		hashedPassword, err := utils.HashPasssword(user.Password)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, "Invalid Password, Failed to create user"+err.Error())
			return
		}

		user.Password = hashedPassword
		err = userService.CreateUser(&user)

		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, "Error while saving user in DB, Failed to create user"+err.Error())
			return
		}
		writeSuccessResponse(w, http.StatusCreated, "User created successfully", user.ID, "")
	}
}

func Login(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var credentials models.User
		err := json.NewDecoder(req.Body).Decode(&credentials)

		if err != nil || !utils.ValidateUserCredentials(&credentials) {
			msg := "Invalid credentials, Failed to login"
			if err != nil {
				msg += ", " + err.Error()
			}
			writeErrorResponse(w, http.StatusBadRequest, msg)
			return
		}

		user, err := userService.GetUserByUsername(credentials.Username)
		if err != nil {
			writeErrorResponse(w, http.StatusNotFound, "User not found, Failed to login, "+err.Error())
			return
		}

		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)) != nil {
			writeErrorResponse(w, http.StatusUnauthorized, "Invalid credentials, Incorrect Password ")
			return
		}

		token, err := utils.GenerateToken(user.ID)

		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, "failed to generate token"+err.Error())
			return
		}

		writeSuccessResponse(w, http.StatusOK, "Successfully Logged In", 0, token)
	}
}

func Logout(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// To log out, we just need to send a success response and have the client remove the token.
		// Assuming the client will delete the JWT from local storage or cookies on the front-end.

		writeSuccessResponse(w, http.StatusOK, "Successfully Logged Out", 0, "")
	}
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	log.Println(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}

func writeSuccessResponse(w http.ResponseWriter, statusCode int, message string, userID uint, token string) {
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

func writeSuccessUserResponse(w http.ResponseWriter, statusCode int, user *models.User) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(user)
}
