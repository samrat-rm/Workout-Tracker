package handlers

import (
	"net/http"
	"workout-tracker/models"
	"workout-tracker/services"
	"workout-tracker/utils"
)

type WorkoutPlan interface {
	CreateWorkoutPlanForUser(w http.ResponseWriter, req *http.Request)
}

type workoutPlan struct {
	userService services.UserService
	workoutPlan services.WorkoutPlanService
}

func NewWokoutSericeHandler(wp services.WorkoutPlanService, us services.UserService) WorkoutPlan {
	return &workoutPlan{
		workoutPlan: wp,
		userService: us,
	}
}

func (wp *workoutPlan) CreateWorkoutPlanForUser(w http.ResponseWriter, req *http.Request) {
	userId, err := fetchUserID(req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "User ID invalid, Please provide a valid user ID, ", err)
		return
	}

	var workoutPlan models.WorkoutPlan
	if err := decodeRequestBody(req, &workoutPlan); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Workout plan data in req body id invalid, Please provide a valid Workout plan data, ", err)
		return
	}

	if err = wp.workoutPlan.CreateWorkoutPlanForUser(userId, &workoutPlan); err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to create workout plan.", err)
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, "Workout Plan successfully added ", nil, nil)
}
