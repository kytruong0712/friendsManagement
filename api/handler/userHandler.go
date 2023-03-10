package handler

import (
	"backend/constants"
	dbrepo "backend/infrastructure/repository/dbRepo"
	"backend/utils"
	"errors"
	"fmt"
	"net/http"

	"github.com/mcnijman/go-emailaddress"
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

	isBlocked, erro := dbrepo.VerifyBlock(email, friend)
	if erro != nil {
		utils.ErrorJSON(w, erro, http.StatusBadRequest)
		return
	}

	if isBlocked.Blocked {
		resp := utils.JSONResponse{
			Success: false,
			Message: fmt.Sprintf("Cannot add friend because %s has blocked %s", email, friend),
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	} else {
		resp := utils.JSONResponse{
			Success: true,
			Message: "create a friend connection successfully",
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	}

}

func CreateSubscribe(w http.ResponseWriter, r *http.Request) {
	// read json payload
	var requestPayload struct {
		Requestor string `json:"requestor"`
		Target    string `json:"target"`
	}

	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	requestor := requestPayload.Requestor
	target := requestPayload.Target

	err = dbrepo.InsertFriend(requestor, target, constants.AddSubscribeToExistingSubscribeArray)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = dbrepo.InsertFriend(requestor, target, constants.AddSubscribeToNullSubscribeArray)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	isBlocked, erro := dbrepo.VerifyBlock(requestor, target)
	if erro != nil {
		utils.ErrorJSON(w, erro, http.StatusBadRequest)
		return
	}

	if isBlocked.Blocked {
		resp := utils.JSONResponse{
			Success: false,
			Message: fmt.Sprintf("Cannot subscribe because %s has blocked %s", requestor, target),
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	} else {
		resp := utils.JSONResponse{
			Success: true,
			Message: "create subscribe successfully",
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	}

}

func CreateBlock(w http.ResponseWriter, r *http.Request) {
	// read json payload
	var requestPayload struct {
		Requestor string `json:"requestor"`
		Target    string `json:"target"`
	}

	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	requestor := requestPayload.Requestor
	target := requestPayload.Target

	err = dbrepo.InsertFriend(requestor, target, constants.AddBlockToExistingSubscribeArray)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = dbrepo.InsertFriend(requestor, target, constants.AddBlockToNullSubscribeArray)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	resp := utils.JSONResponse{
		Success: true,
		Message: "connection was blocked successfully",
	}

	utils.WriteJSON(w, http.StatusOK, resp)
}

func RetrieveUpdates(w http.ResponseWriter, r *http.Request) {
	// read json payload
	var requestPayload struct {
		Sender string `json:"sender"`
		Text   string `json:"text"`
	}

	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	sender := requestPayload.Sender
	mentions := emailaddress.Find([]byte(requestPayload.Text), false)

	users, err := dbrepo.GetUser(sender)
	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	rereiveList := make([]string, 0)
	rereiveList = utils.AppendWithoutDuplicate(rereiveList, users.Friends)
	rereiveList = utils.AppendWithoutDuplicate(rereiveList, users.Subscribe)
	for _, m := range mentions {
		rereiveList = utils.AppendWithoutDuplicate(rereiveList, []string{m.LocalPart + "@" + m.Domain})
	}

	rereiveList = utils.FindMissing(rereiveList, users.Blocks)

	resp := utils.JSONReceiveUpdates{
		Success:    true,
		Message:    "retreive updates successfully",
		Recipients: rereiveList,
	}

	utils.WriteJSON(w, http.StatusOK, resp)

}
