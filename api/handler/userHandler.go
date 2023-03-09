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
	err := dbrepo.InsertFriend("andrew@example.com", "donald@example.com",
		`update public.user
		set    friends = (select array_agg(distinct e) from unnest(friends || ARRAY[$2]) e)
		where  not friends @> ARRAY[$2] and email = $1;`)

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	err = dbrepo.InsertFriend("andrew@example.com", "donald@example.com",
		`UPDATE public.user SET friends = ARRAY[$2]
		where  email = $1 and friends IS NULL;`)

	if err != nil {
		utils.ErrorJSON(w, err)
		return
	}

	users, _ := dbrepo.AllUsers()

	_ = utils.WriteJSON(w, http.StatusOK, users)
}
