package routing

import (
	"backend/api/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()

	mux.Get("/users", handler.AllUsers)
	mux.Get("/user", handler.GetUser)
	mux.Post("/invite", handler.InsertFriend)
	mux.Post("/friends", handler.GetFriendList)
	mux.Post("/common", handler.GetCommonFriends)
	mux.Post("/subscribe", handler.CreateSubscribe)

	return mux
}
