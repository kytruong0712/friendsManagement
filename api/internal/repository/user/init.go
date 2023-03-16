package repository

import (
	"database/sql"
)

// UserRepository: User Repository
type UserRepository struct {
	db *sql.DB
}

// NewUserRepo: create new User repository
func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
