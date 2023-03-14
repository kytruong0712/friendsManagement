package main

import (
	"backend/api/internal/app/connection"
	"backend/api/internal/config"
	"backend/api/internal/handler/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5",
		config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASSWORD, config.DB_DATABASE)

	// connect to the database
	err := connection.ConnectToDB(dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer connection.DBConn.Close()

	log.Println("Starting application on port", config.API_PORT)

	// start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.API_PORT), router.Routes())
	if err != nil {
		log.Fatal(err)
	}

}
