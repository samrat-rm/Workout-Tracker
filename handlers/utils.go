package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func decodeRequestBody(req *http.Request, v interface{}) error {
	return json.NewDecoder(req.Body).Decode(v)
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

func fetchWorkoutPlanID(req *http.Request) (uint, error) {
	vars := mux.Vars(req)
	id := vars["wp_id"]
	userId, err := strconv.Atoi(id)

	if err != nil {
		fmt.Printf("Error while converting userId to uint %s", id)
		return 0, err
	}
	return uint(userId), nil
}
