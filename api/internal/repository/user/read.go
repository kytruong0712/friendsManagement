package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"

	"backend/api/internal/models"
	"backend/api/internal/presenter"
	"backend/api/pkg/constants"
)

const dbTimeout = time.Second * 3

// List: Get All users from database
func (repo *UserRepository) List() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := constants.GetAllUsers

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			pq.Array(&user.Friends),
			pq.Array(&user.Subscribe),
			pq.Array(&user.Blocks),
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Get: Get a single user by email
func (repo *UserRepository) Get(email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := constants.GetUser

	user := models.User{}

	row := repo.db.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		pq.Array(&user.Friends),
		pq.Array(&user.Subscribe),
		pq.Array(&user.Blocks),
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return user, err
	}

	return user, nil

}

// Verify whether requestor has already blocked target or not
func (repo *UserRepository) IsBlock(requestor string, target string) (*presenter.IsBlock, error) {
	if requestor == target {
		return nil, errors.New("2 input emails are the same")
	}

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	_, err := repo.Get(requestor)
	if err != nil {
		return nil, err
	}

	_, errTarget := repo.Get(target)
	if errTarget != nil {
		return nil, errTarget
	}

	query := constants.VerifyBlock

	row := repo.db.QueryRowContext(ctx, query, requestor, target)

	var blocked presenter.IsBlock

	err = row.Scan(
		&blocked.Blocked,
	)

	if err != nil {
		fmt.Println("blocked error", err)
		return nil, err
	}

	return &blocked, nil
}