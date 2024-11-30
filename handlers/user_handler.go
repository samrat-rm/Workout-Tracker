package handlers

import (
	"net/http"
	"workout-tracker/models"
	"workout-tracker/services"
	"workout-tracker/utils"
)

func GetUser(userService services.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		userId, err := fetchUserID(req)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "User ID invalid, Please provide a valid user ID, ", err)
			return
		}
		user, err := userService.GetUser(userId)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusNotFound, "User not found, ", err)
			return
		}

		utils.WriteSuccessUserResponse(w, http.StatusOK, user)
	}
}

func DeleteUser(userService services.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		userId, err := fetchUserID(req)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "User ID invalid, Please provide a valid user ID, ", err)
			return
		}
		err = userService.DeleteUser(uint(userId))
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusNotFound, "User not found, ", err)
			return
		}

		utils.WriteSuccessResponse(w, http.StatusNoContent, "User deleted successfully", &userId, nil)
	}
}

func UpdateUser(userService services.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		userId, err := fetchUserID(req)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "User ID invalid, Please provide a valid user ID, ", err)
			return
		}

		var user models.User
		if err := decodeRequestBody(req, user); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "User data in req body id invalid, Please provide a valid user data, ", err)
			return
		}

		err = userService.UpdateUser(uint(userId), &user)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusNotFound, "User not found, ", err)
			return
		}

		utils.WriteSuccessResponse(w, http.StatusOK, "User Updated successfully", &userId, nil)
	}
}

func AddWorkoutToUser(userService services.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		userID, err := fetchUserID(req)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "User ID invalid, Please provide a valid user ID, ", err)
			return
		}

		var workoutPlan models.WorkoutPlan
		if err := decodeRequestBody(req, workoutPlan); err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, "WorkoutPlan data in req body id invalid, Please provide a valid workoutPlan data, ", err)
			return
		}

		err = userService.AddWorkoutToUser(userID, &workoutPlan)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, "Error while creating workoutPlan for user, ", err)
			return
		}
		utils.WriteSuccessResponse(w, http.StatusCreated, "Workout plan created successfully", &userID, nil)

	}
}
