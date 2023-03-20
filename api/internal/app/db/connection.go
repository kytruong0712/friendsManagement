package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	pkgErr "github.com/pkg/errors"
	"os"
	"fmt"
)

// Connect: uses to connect to the database
func Connect(dns string) (*sql.DB, error) {
	conn, err := openDB(dns)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to Postgres!")
	return conn, nil
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)

		return nil, pkgErr.WithStack(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping to database: %v\n", err)
		return nil, pkgErr.WithStack(err)
	}

	fmt.Fprintln(os.Stderr, "connect to DB successfully")

	return db, nil
}
