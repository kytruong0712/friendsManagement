package handler

import (
	dbrepo "backend/infrastructure/repository/dbRepo"
	"backend/utils"
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
	err := dbrepo.InsertFriend("andrew@example.com", "donald@example.com")

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	users, _ := dbrepo.AllUsers()

	_ = utils.WriteJSON(w, http.StatusOK, users)
}
