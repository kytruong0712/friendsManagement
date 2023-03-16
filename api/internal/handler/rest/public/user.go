package handler

import (
	"backend/api/pkg/utils"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mcnijman/go-emailaddress"

	controller "backend/api/internal/controller/user"
)

func list(controller controller.UserInterface) (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users, err := controller.List()
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, users)
	})
}

func get(controller controller.UserInterface) (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		users, err := controller.Get(email)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, users)
	})
}

func createFriendship(controller controller.UserInterface) (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		resp, er := controller.CreateFriendship(email, friend)
		if er != nil {
			utils.ErrorJSON(w, er)
			return
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	})
}

func getFriendList(controller controller.UserInterface) (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		friendList, er := controller.GetFriendList(email)
		if er != nil {
			utils.ErrorJSON(w, er)
			return
		}

		utils.WriteJSON(w, http.StatusOK, friendList)
	})
}

func getCommonFriends(controller controller.UserInterface) (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		friendList, er := controller.GetCommonFriends(email, friend)
		if er != nil {
			utils.ErrorJSON(w, er)
			return
		}

		utils.WriteJSON(w, http.StatusOK, friendList)
	})
}

func createSubscribe(controller controller.UserInterface) (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		resp, er := controller.CreateSubscribe(requestor, target)
		if er != nil {
			utils.ErrorJSON(w, er)
			return
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	})
}

func createBlock(controller controller.UserInterface) (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		err = controller.CreateBlock(requestor, target)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		resp := utils.JSONResponse{
			Success: true,
			Message: "connection was blocked successfully",
		}

		utils.WriteJSON(w, http.StatusOK, resp)
	})
}

func getRetrieveUpdates(controller controller.UserInterface) (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		retrieveUpdatesResp, er := controller.GetRetrieveUpdates(sender, mentions)

		if er != nil {
			utils.ErrorJSON(w, er)
			return
		}

		utils.WriteJSON(w, http.StatusOK, retrieveUpdatesResp)
	})
}

// MakeUserHandlers: make url handlers
func MakeUserHandlers(mux *chi.Mux, controller controller.UserInterface) http.Handler {

	mux.Get("/users", list(controller))
	mux.Post("/user", get(controller))
	mux.Post("/invite", createFriendship(controller))
	mux.Post("/friends", getFriendList(controller))
	mux.Post("/common", getCommonFriends(controller))
	mux.Post("/subscribe", createSubscribe(controller))
	mux.Post("/blocks", createBlock(controller))
	mux.Post("/retrieve", getRetrieveUpdates(controller))

	return mux
}
