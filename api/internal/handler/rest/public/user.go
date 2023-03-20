package user

import (
	"errors"
	"net/http"

	"github.com/mcnijman/go-emailaddress"

	controller "backend/api/internal/controller/user"
	"backend/api/internal/presenter"
	"backend/api/pkg/utils"
)

// read json payload
type emailRequestPayload struct {
	Email string `json:"email"`
}

type friendRequestPayload struct {
	Friends []string `json:"friends"`
}

type requestorRequestPayload struct {
	Requestor string `json:"requestor"`
	Target    string `json:"target"`
}

type senderRequestPayload struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

// TODO: create type in this file instead of inside function

func List(controller controller.UserInterface) (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := controller.List()
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}
		var users []presenter.User
		for _, d := range data {
			users = append(users, presenter.User{
				ID:        d.ID,
				Name:      d.Name,
				Email:     d.Email,
				Friends:   d.Friends,
				Subscribe: d.Subscribe,
				Blocks:    d.Blocks,
				CreatedAt: d.CreatedAt,
				UpdatedAt: d.UpdatedAt,
			})
		}

		utils.WriteJSON(w, http.StatusOK, users)
	})
}

func Get(controller controller.UserInterface) (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPayload := emailRequestPayload{}
		err := utils.ReadJSON(w, r, &requestPayload)
		if err != nil {
			utils.ErrorJSON(w, err, http.StatusBadRequest)
			return
		}

		users, err := controller.Get(requestPayload.Email)
		if err != nil {
			utils.ErrorJSON(w, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, users)
	})
}

func CreateFriendship(controller controller.UserInterface) (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPayload := friendRequestPayload{}

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

func GetFriendList(controller controller.UserInterface) (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPayload := emailRequestPayload{}

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

func GetCommonFriends(controller controller.UserInterface) (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPayload := friendRequestPayload{}

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

func CreateSubscribe(controller controller.UserInterface) (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPayload := requestorRequestPayload{}

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

func CreateBlock(controller controller.UserInterface) (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPayload := requestorRequestPayload{}

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

func GetRetrieveUpdates(controller controller.UserInterface) (handlerFn http.HandlerFunc) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestPayload := senderRequestPayload{}

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
