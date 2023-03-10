package dbrepo

import (
	"backend/api/presenter"
	"backend/constants"
	"backend/utils"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
)

const dbTimeout = time.Second * 3

func AllUsers() ([]*presenter.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := constants.GetAllUsers

	rows, err := utils.DBConn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*presenter.User

	for rows.Next() {
		var user presenter.User
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

		users = append(users, &user)
	}

	return users, nil
}

func GetUser(email string) (*presenter.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := constants.GetUser

	row := utils.DBConn.QueryRowContext(ctx, query, email)

	var user presenter.User

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
		return nil, err
	}

	return &user, nil
}

func InsertFriend(email string, friend string, stmt string) error {
	if email == friend {
		return errors.New("2 input emails are the same")
	}

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	_, err := GetUser(email)
	if err != nil {
		return err
	}

	_, errFriend := GetUser(friend)
	if errFriend != nil {
		return errFriend
	}

	_, erro := utils.DBConn.ExecContext(ctx, stmt,
		email,
		friend,
	)
	if erro != nil {
		fmt.Println("erro", erro)
		return err
	}

	return nil
}

func VerifyBlock(requestor string, target string) (*presenter.IsBlock, error) {
	if requestor == target {
		return nil, errors.New("2 input emails are the same")
	}

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	_, err := GetUser(requestor)
	if err != nil {
		return nil, err
	}

	_, errTarget := GetUser(target)
	if errTarget != nil {
		return nil, errTarget
	}

	query := constants.VerifyBlock

	row := utils.DBConn.QueryRowContext(ctx, query, requestor, target)

	var blocked presenter.IsBlock

	err = row.Scan(
		&blocked.Blocked,
	)

	if err != nil {
		return nil, err
	}

	return &blocked, nil
}
