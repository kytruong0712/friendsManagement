package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	controller "backend/api/internal/controller/user"
	handler "backend/api/internal/handler/rest/public"
)

// Routes: Router of users
func Routes(router *chi.Mux, controller controller.UserInterface) http.Handler {
	router.Get("/users", handler.List(controller))
	router.Post("/user", handler.Get(controller))
	router.Post("/invite", handler.CreateFriendship(controller))
	router.Post("/friends", handler.GetFriendList(controller))
	router.Post("/common", handler.GetCommonFriends(controller))
	router.Post("/subscribe", handler.CreateSubscribe(controller))
	router.Post("/blocks", handler.CreateBlock(controller))
	router.Post("/retrieve", handler.GetRetrieveUpdates(controller))

	return router
}
