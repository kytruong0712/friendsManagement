package mod

type UserResponse struct {
	Success bool
	Message string
}

// IsBlock: mod for checking whether a requestor has blocked target or not
type IsBlock struct {
	Blocked bool
}

// FriendList: mod for friend list
type FriendList struct {
	Success bool
	Friends []string
	Count   int
}

// RetrieveUpdates: mod for retrieving updates
type RetrieveUpdates struct {
	Success    bool
	Message    string
	Recipients []string
}
