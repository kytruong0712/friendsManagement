package presenter

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Friends   []string  `json:"friends"`
	Subscribe []string  `json:"subscribe"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
