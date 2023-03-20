package main

import (
	"log"
	"net/http"
	"os"

	pkgErr "github.com/pkg/errors"

	"gihub.com/AI-LastWish/friendsManagement/api/internal/app/db"
	controller "gihub.com/AI-LastWish/friendsManagement/api/internal/controller/user"
	repository "gihub.com/AI-LastWish/friendsManagement/api/internal/repository/user"
	"gihub.com/AI-LastWish/friendsManagement/api/pkg/constants"
	"fmt"
	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("Holla!")
	dataSourceName := "host=database user=postgres password=postgres dbname=backend port=5432 sslmode=disable"

	// connect to the database
	conn, err := db.Connect(dataSourceName)
	if err != nil {
		log.Fatal("Err: ", err.Error())
	}
	defer conn.Close()

	// create a router mux
	router := chi.NewRouter()

	userRepo := repository.NewUserRepo(conn)
	userController := controller.NewUserController(userRepo)

	log.Println("Starting application on port: ", os.Getenv(constants.API_PORT))

	// start a web server
	if err = http.ListenAndServe(":3003", routes(router, userController)); err != nil {
		log.Fatal("Err: ", pkgErr.WithStack(err))
	}
}
