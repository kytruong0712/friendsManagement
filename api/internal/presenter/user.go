package presenter

import "time"

// User: responsible for formatting the data generated as a response by handlers
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Friends   []string  `json:"friends,omitempty"`
	Subscribe []string  `json:"subscribe,omitempty"`
	Blocks    []string  `json:"blocks,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type IsBlock struct {
	Blocked bool `json:"blocked"`
}
