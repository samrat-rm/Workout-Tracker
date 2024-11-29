package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"workout-tracker/models"
	"workout-tracker/services"
	"workout-tracker/utils"

	"github.com/gorilla/mux"
)

func GetUser(userService services.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		userId, err := fetchUserID(req)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "User ID invalid, Please provide a valid user ID, "+err.Error())
			return
		}
		user, err := userService.GetUser(userId)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusNotFound, "User not found, "+err.Error())
			return
		}

		if err := utils.WriteSuccessUserResponse(w, http.StatusOK, user); err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to encode user data "+err.Error())
			return
		}

	}
}

func DeleteUser(userService services.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		userId, err := fetchUserID(req)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "User ID invalid, Please provide a valid user ID, "+err.Error())
			return
		}
		err = userService.DeleteUser(uint(userId))
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusNotFound, "User not found, "+err.Error())
			return
		}

		utils.WriteSuccessResponse(w, http.StatusNoContent, "User deleted successfully", userId, "")
	}
}

func UpdateUser(userService services.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		var user models.User
		userId, err := fetchUserID(req)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "User ID invalid, Please provide a valid user ID, "+err.Error())
			return
		}

		if err = json.NewDecoder(req.Body).Decode(&user); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "User data in req body id invalid, Please provide a valid user data, "+err.Error())
			return
		}

		err = userService.UpdateUser(uint(userId), &user)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusNotFound, "User not found, "+err.Error())
			return
		}

		utils.WriteSuccessResponse(w, http.StatusOK, "User Updated successfully", userId, "")
	}
}

func AddWorkoutToUser(userService services.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		var workoutPlan models.WorkoutPlan
		userID, err := fetchUserID(req)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "User ID invalid, Please provide a valid user ID, "+err.Error())
			return
		}

		if err = json.NewDecoder(req.Body).Decode(&workoutPlan); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "WorkoutPlan data in req body id invalid, Please provide a valid workoutPlan data, "+err.Error())
			return
		}

		err = userService.AddWorkoutToUser(userID, &workoutPlan)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, "Error while creating workoutPlan for user, "+err.Error())
			return
		}
		utils.WriteSuccessResponse(w, http.StatusCreated, "Workout plan created successfully", userID, "")

	}
}

func fetchUserID(req *http.Request) (uint, error) {
	vars := mux.Vars(req)
	id := vars["id"]
	userId, err := strconv.Atoi(id)

	if err != nil {
		fmt.Printf("Error while converting userId to uint %s", id)
		return 0, err
	}
	return uint(userId), nil
}
