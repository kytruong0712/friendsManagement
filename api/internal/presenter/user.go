package presenter

import "time"

// User: responsible for formatting User generated as a response
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

// IsBlock: responsible for checking whether a requestor has blocked target or not
type IsBlock struct {
	Blocked bool `json:"blocked"`
}

// FriendList: responsible for formatting list of friends as a response
type FriendList struct {
	Success bool     `json:"success"`
	Friends []string `json:"friends"`
	Count   int      `json:"count"`
}

// RetrieveUpdates: responsible for formatting retrieving updates as a response
type RetrieveUpdates struct {
	Success    bool     `json:"success"`
	Message    string   `json:"message"`
	Recipients []string `json:"recipients"`
}
