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

func GetFriendList(w http.ResponseWriter, r *http.Request) {
	// read json payload
	var requestPayload struct {
		Email string `json:"email"`
	}

	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	email := requestPayload.Email

	users, err := dbrepo.GetUser(email)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	count := len(users.Friends)

	friendsList := make([]string, 0)
	if count > 0 {
		friendsList = users.Friends
	}

	resp := utils.JSONFriendList{
		Success: true,
		Friends: friendsList,
		Count:   count,
	}

	utils.WriteJSON(w, http.StatusOK, resp)
}

func GetCommonFriends(w http.ResponseWriter, r *http.Request) {
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

	users1, err1 := dbrepo.GetUser(email)
	if err1 != nil {
		utils.ErrorJSON(w, err1)
		return
	}

	users2, err2 := dbrepo.GetUser(friend)
	if err2 != nil {
		utils.ErrorJSON(w, err2)
		return
	}

	friends1 := make([]string, 0)
	if len(users1.Friends) > 0 {
		friends1 = users1.Friends
	}

	friends2 := make([]string, 0)
	if len(users2.Friends) > 0 {
		friends2 = users2.Friends
	}

	temp_intersect := utils.HashGeneric(friends1, friends2)
	intersect := make([]string, 0)
	for _, value := range temp_intersect {
		if value != email && value != friend {
			intersect = append(intersect, value)
		}
	}

	resp := utils.JSONFriendList{
		Success: true,
		Friends: intersect,
		Count:   len(intersect),
	}

	utils.WriteJSON(w, http.StatusOK, resp)

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
