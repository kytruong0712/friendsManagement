package repository

import (
	"backend/api/internal/presenter"
	"database/sql"
)

// UserRepository: repository of User entity
type UserRepository interface {
	Connection() *sql.DB
	AllUsers() ([]*presenter.User, error)
}
