package repository

import (
	"backend/api/presenter"
	"database/sql"
)

type UserRepository interface {
	Connection() *sql.DB
	AllUsers() ([]*presenter.User, error)
}
