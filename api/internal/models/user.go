package models

import "time"

// User: responsible for represent User entity
type User struct {
	ID        int
	Name      string
	Email     string
	Friends   []string
	Subscribe []string
	Blocks    []string
	CreatedAt time.Time
	UpdatedAt time.Time
}
