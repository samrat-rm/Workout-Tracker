package handlers

import (
	"net/http"
	"workout-tracker/models"
	"workout-tracker/services"
	"workout-tracker/utils"
)

type ExerciseHandler interface {
	CreateExercise(w http.ResponseWriter, req *http.Request)
	UpdateExercise(w http.ResponseWriter, req *http.Request)
	DeleteExercise(w http.ResponseWriter, req *http.Request)
}

type exerciseHandler struct {
	exerciseService services.ExerciseService
}

func NewExerciseHandler(es services.ExerciseService) ExerciseHandler {
	return &exerciseHandler{exerciseService: es}
}

func (h *exerciseHandler) CreateExercise(w http.ResponseWriter, req *http.Request) {
	_, err := fetchUserID(req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "User ID invalid, Please provide a valid user ID.", err)
		return
	}
	// VALIDATOR for input

	workoutPlanID, err := fetchWorkoutPlanID(req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Workout Plan ID invalid, Please provide a valid workout plan ID.", err)
		return
	}

	var exercise models.Exercise
	if err := decodeRequestBody(req, &exercise); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid exercise data in request body.", err)
		return
	}

	err = h.exerciseService.CreateExercise(workoutPlanID, exercise)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to create exercise.", err)
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, "Exercise successfully created", nil, nil)
}

func (h *exerciseHandler) UpdateExercise(w http.ResponseWriter, req *http.Request) {
	exerciseID, err := fetchExerciseID(req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Exercise ID invalid, Please provide a valid exercise ID.", err)
		return
	}

	var updatedExercise models.Exercise
	if err := decodeRequestBody(req, &updatedExercise); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid exercise data in request body.", err)
		return
	}

	_, err = h.exerciseService.UpdateExercise(exerciseID, updatedExercise)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to update exercise.", err)
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, "Exercise successfully updated", nil, nil)
}

func (h *exerciseHandler) DeleteExercise(w http.ResponseWriter, req *http.Request) {
	exerciseID, err := fetchExerciseID(req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Exercise ID invalid, Please provide a valid exercise ID.", err)
		return
	}

	if err := h.exerciseService.DeleteExercise(exerciseID); err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to delete exercise.", err)
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, "Exercise successfully deleted", nil, nil)
}
