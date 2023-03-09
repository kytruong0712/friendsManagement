package dbrepo

import (
	"backend/api/presenter"
	"backend/utils"
	"context"
	"time"
)

const dbTimeout = time.Second * 3

func AllUsers() ([]*presenter.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select
			id, name, email, created_at, updated_at
		from
			public.user
		order by
		name
	`

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
