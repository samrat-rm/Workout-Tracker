package handler

import (
	"encoding/json"
	"net/http"
	"workout-tracker/models"
	"workout-tracker/utils"

	"golang.org/x/crypto/bcrypt"
)

func SignUp(w http.ResponseWriter, req *http.Request) {
	var user models.User
	json.NewDecoder(req.Body).Decode(&user)
	user.HasWorkoutPlan = false
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := utils.DB.Create(&user).Error; err != nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, req *http.Request) {
	var credentials models.User
	json.NewDecoder(req.Body).Decode(&credentials)

	var user models.User
	utils.DB.Where("username = ?", credentials.Username).First(&user)

	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateToken(user.ID)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func Logout(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Logged out successfully"}`))
}
