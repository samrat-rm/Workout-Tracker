package handlers

import (
	"net/http"
	"workout-tracker/models"
	"workout-tracker/services"
	"workout-tracker/utils"

	"golang.org/x/crypto/bcrypt"
)

func SignUp(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var user models.User
		if err := decodeRequestBody(req, user); err != nil || !utils.ValidateUserCredentials(&user) {
			msg := "Invalid credentials, Failed to Sign up"
			if err != nil {
				msg += ", " + err.Error()
			}
			utils.WriteErrorResponse(w, http.StatusBadRequest, msg, err)
			return
		}

		hashedPassword, err := utils.HashPasssword(user.Password)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, "Invalid Password, Failed to create user", err)
			return
		}

		user.Password = hashedPassword
		err = userService.CreateUser(&user)

		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, "Error while saving user in DB, Failed to create user", err)
			return
		}
		utils.WriteSuccessResponse(w, http.StatusCreated, "User created successfully", &user.ID, nil)
	}
}

func Login(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		var credentials models.User
		if err := decodeRequestBody(req, credentials); err != nil || !utils.ValidateUserCredentials(&credentials) {
			msg := "Invalid credentials, Failed to login"
			if err != nil {
				msg += ", " + err.Error()
			}
			utils.WriteErrorResponse(w, http.StatusBadRequest, msg, err)
			return
		}

		user, err := userService.GetUserByUsername(credentials.Username)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusNotFound, "User not found, Failed to login, ", err)
			return
		}

		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)) != nil {
			utils.WriteErrorResponse(w, http.StatusUnauthorized, "Invalid credentials, Incorrect Password ", err)
			return
		}

		token, err := utils.GenerateToken(user.ID)

		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, "failed to generate token", err)
			return
		}

		utils.WriteSuccessResponse(w, http.StatusOK, "Successfully Logged In", nil, &token)
	}
}

func Logout(userService services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// To log out, we just need to send a success response and have the client remove the token.
		// Assuming the client will delete the JWT from local storage or cookies on the front-end.

		utils.WriteSuccessResponse(w, http.StatusOK, "Successfully Logged Out", nil, nil)
	}
}
