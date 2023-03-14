package models

type User struct {
	ID        int
	Name      string
	Email     string
	Friends   []string
	Subscribe []string
}
