package handler

import (
	"backend/constants"
	dbrepo "backend/infrastructure/repository/dbRepo"
	"backend/utils"
	"errors"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Movies up and running",
		Version: "1.0.0",
	}

	_ = utils.WriteJSON(w, http.StatusOK, payload)
}

func AllUsers(w http.ResponseWriter, r *http.Request) {

	users, err := dbrepo.AllUsers()
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	_ = utils.WriteJSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	users, err := dbrepo.GetUser("tom@example.com")
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	_ = utils.WriteJSON(w, http.StatusOK, users)
}

func InsertFriend(w http.ResponseWriter, r *http.Request) {
	// read json payload
	var requestPayload struct {
		Friends []string `json:"friends"`
	}

	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	if len(requestPayload.Friends) != 2 {
		utils.ErrorJSON(w, errors.New("invalid input"), http.StatusBadRequest)
	}

	email := requestPayload.Friends[0]
	friend := requestPayload.Friends[1]

	err = dbrepo.InsertFriend(email, friend, constants.AddFriendToExistingFriendsArray)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = dbrepo.InsertFriend(email, friend, constants.AddFriendToNullFriendsArray)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = dbrepo.InsertFriend(friend, email, constants.AddFriendToExistingFriendsArray)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = dbrepo.InsertFriend(friend, email, constants.AddFriendToNullFriendsArray)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp := utils.JSONResponse{
		Success: true,
		Message: "create a friend connection successfully",
	}

	utils.WriteJSON(w, http.StatusOK, resp)
}
