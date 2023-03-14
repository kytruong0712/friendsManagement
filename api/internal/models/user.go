package models

// User: responsible for represent User entity
type User struct {
	ID        int
	Name      string
	Email     string
	Friends   []string
	Subscribe []string
}
