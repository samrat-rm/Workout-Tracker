package handlers

import (
	"net/http"
	"workout-tracker/models"
	"workout-tracker/services"
	"workout-tracker/utils"
)

type UserHandler interface {
	GetUser(w http.ResponseWriter, req *http.Request)
	DeleteUser(w http.ResponseWriter, req *http.Request)
	UpdateUser(w http.ResponseWriter, req *http.Request)
}

type userHandler struct {
	userService services.UserService
}

func NewUserHandler(us services.UserService) UserHandler {
	return &userHandler{userService: us}
}

func (h *userHandler) GetUser(w http.ResponseWriter, req *http.Request) {
	userId, err := fetchUserID(req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "User ID invalid, Please provide a valid user ID, ", err)
		return
	}
	user, err := h.userService.GetUser(userId)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, "User not found, ", err)
		return
	}

	utils.WriteSuccessResponseWithBody(w, http.StatusOK, user)
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, req *http.Request) {
	userId, err := fetchUserID(req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "User ID invalid, Please provide a valid user ID, ", err)
		return
	}
	err = h.userService.DeleteUser(uint(userId))
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, "User not found, ", err)
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, "User deleted successfully", &userId, nil)
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, req *http.Request) {
	userId, err := fetchUserID(req)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "User ID invalid, Please provide a valid user ID, ", err)
		return
	}

	var user models.User
	if err := decodeRequestBody(req, &user); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "User data in req body id invalid, Please provide a valid user data, ", err)
		return
	}

	err = h.userService.UpdateUser(uint(userId), &user)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, "User not found, ", err)
		return
	}

	utils.WriteSuccessResponse(w, http.StatusOK, "User Updated successfully", &userId, nil)
}
