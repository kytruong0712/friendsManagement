package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"backend/api/internal/app/db"
	"backend/api/internal/config"
	controller "backend/api/internal/controller/user"
	handler "backend/api/internal/handler/rest/public"
	repository "backend/api/internal/repository/user"

	pkgErr "github.com/pkg/errors"
)

func main() {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5",
		config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASSWORD, config.DB_DATABASE)

	// connect to the database
	conn, err := db.Connect(dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	// create a router mux
	router := chi.NewRouter()

	userRepo := repository.NewUserRepo(conn)
	userController := controller.NewUserController(userRepo)

	log.Println("Starting application on port", config.API_PORT)

	// start a web server
	if err = http.ListenAndServe(fmt.Sprintf(":%d", config.API_PORT), handler.MakeUserHandlers(router, userController)); err != nil {
		log.Fatal(pkgErr.WithStack(err))
	}
}
