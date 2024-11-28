package handlers

import (
	"net/http"
	"strconv"
	"workout-tracker/services"

	"github.com/gorilla/mux"
)

func GetUser(userService services.UserService) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		id := vars["id"]

		userId, err := strconv.Atoi(id)
		if err != nil {
			writeErrorResponse(w, http.StatusBadRequest, "User ID invalid, Please provide a valid user ID, "+err.Error())
			return
		}
		user, err := userService.GetUser(uint(userId))
		if err != nil {
			writeErrorResponse(w, http.StatusNotFound, "User not found, "+err.Error())
			return
		}

		if err := writeSuccessUserResponse(w, http.StatusOK, user); err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, "Failed to encode user data "+err.Error())
			return
		}

	}
}
