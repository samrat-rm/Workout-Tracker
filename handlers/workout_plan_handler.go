package handlers

import (
	"net/http"
	"workout-tracker/models"
	"workout-tracker/services"
	"workout-tracker/utils"
)

type WorkoutPlan interface {
	CreateWorkoutPlanForUser(w http.ResponseWriter, req *http.Request)
	GetAllWorkoutPlansForUser(w http.ResponseWriter, req *http.Request)
	GetWorkoutPlanForUser(w http.ResponseWriter, req *http.Request)
	UpdateWorkoutPlanForUser(w http.ResponseWriter, req *http.Request)
	RemoveWorkoutPlanForUser(w http.ResponseWriter, req *http.Request)
	UpdateWorkoutPlanStatusForUser(w http.ResponseWriter, req *http.Request)
}

type workoutPlan struct {
	userService services.UserService
	workoutPlan services.WorkoutPlanService
}

func NewWorkoutServiceHandler(wp services.WorkoutPlanService, us services.UserService) WorkoutPlan {
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

func (wp *workoutPlan) GetAllWorkoutPlansForUser(w http.ResponseWriter, req *http.Request) {
	userId, err := fetchUserID(req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "User ID invalid, Please provide a valid user ID, ", err)
		return
	}

	workoutPlans, err := wp.workoutPlan.GetAllWorkoutPlansForUser(userId)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, "Error while finding workout plans, ", err)
		return
	}

	utils.WriteSuccessResponseWithBody(w, http.StatusOK, workoutPlans)
}

func (wp *workoutPlan) GetWorkoutPlanForUser(w http.ResponseWriter, req *http.Request) {
	userID, err := fetchUserID(req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "user ID invalid, Please provide a valid user ID, ", err)
		return
	}

	wpID, err := fetchWorkoutPlanID(req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "ID invalid, Please provide a valid workout plan ID, ", err)
		return
	}

	workoutPlan, err := wp.workoutPlan.GetWorkoutPlanForUser(userID, wpID)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, "Error while finding workout plan, ", err)
		return
	}

	utils.WriteSuccessResponseWithBody(w, http.StatusOK, workoutPlan)
}

func (wp *workoutPlan) UpdateWorkoutPlanForUser(w http.ResponseWriter, req *http.Request) {
	userID, err := fetchUserID(req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "user ID invalid, Please provide a valid user ID, ", err)
		return
	}

	wp_id, err := fetchWorkoutPlanID(req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "ID invalid, Please provide a valid workout plan ID, ", err)
		return
	}

	var updatedWorkoutPlan models.WorkoutPlan
	if err := decodeRequestBody(req, &updatedWorkoutPlan); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid workout plan data in request body, ", err)
		return
	}

	if err := wp.workoutPlan.UpdateWorkoutPlanforUser(userID, wp_id, &updatedWorkoutPlan); err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to update workout plan.", err)
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, "Workout Plan successfully updated", nil, nil)
}

func (wp *workoutPlan) RemoveWorkoutPlanForUser(w http.ResponseWriter, req *http.Request) {
	userID, err := fetchUserID(req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "user ID invalid, Please provide a valid user ID, ", err)
		return
	}

	wpID, err := fetchWorkoutPlanID(req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "ID invalid, Please provide a valid workout plan ID, ", err)
		return
	}

	if err := wp.workoutPlan.RemoveWorkoutPlanForUser(userID, wpID); err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to remove workout plan.", err)
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, "Workout Plan successfully removed", nil, nil)
}

func (wp *workoutPlan) UpdateWorkoutPlanStatusForUser(w http.ResponseWriter, req *http.Request) {
	// Step 1: Fetch user ID from the request (you might already have a helper function for this)
	userID, err := fetchUserID(req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "User ID invalid, Please provide a valid user ID", err)
		return
	}

	// Step 2: Fetch workout plan ID from the request (again, a helper function could be useful here)
	wpID, err := fetchWorkoutPlanID(req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Workout Plan ID invalid, Please provide a valid workout plan ID", err)
		return
	}

	// Step 3: Decode the request body into a map to extract the status field
	var requestBody map[string]interface{}
	if err := decodeRequestBody(req, &requestBody); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid data in request body", err)
		return
	}

	// Step 4: Ensure status is provided and convert it to the correct type
	statusStr, ok := requestBody["status"].(string)
	if !ok {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Status field is required and should be a string", nil)
		return
	}

	var status models.Status
	// Attempt to unmarshal the status string into the correct Status enum value
	if err := status.UnmarshalJSON([]byte(`"` + statusStr + `"`)); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid status value", err)
		return
	}

	// Step 5: Update the workout plan status using the service method
	if err := wp.workoutPlan.UpdateWorkoutPlanStatusForUser(userID, wpID, status); err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to update workout plan status", err)
		return
	}

	// Step 6: Send a success response
	utils.WriteSuccessResponse(w, http.StatusOK, "Workout Plan status successfully updated", nil, nil)
}
