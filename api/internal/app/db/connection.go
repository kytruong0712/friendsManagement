package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// ConnectToDB: use to connect to the database
func ConnectToDB(dns string) (*sql.DB, error) {
	conn, err := openDB(dns)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Postgres!")
	return conn, nil
}
