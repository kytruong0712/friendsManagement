package repository

import (
	"backend/api/internal/presenter"
	"database/sql"
)

type UserRepository interface {
	Connection() *sql.DB
	AllUsers() ([]*presenter.User, error)
}
