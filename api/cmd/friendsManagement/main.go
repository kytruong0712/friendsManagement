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
)

func main() {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5",
		config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASSWORD, config.DB_DATABASE)

	// connect to the database
	conn, err := db.ConnectToDB(dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	// create a router mux
	mux := chi.NewRouter()

	userRepo := repository.NewUserRepo(conn)
	userController := controller.NewUserController(userRepo)

	log.Println("Starting application on port", config.API_PORT)

	// start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.API_PORT), handler.MakeUserHandlers(mux, userController))
	if err != nil {
		log.Fatal(err)
	}

}
