package user

import (
	"github.com/mcnijman/go-emailaddress"

	"backend/api/internal/mod"
	"backend/api/internal/models"
	"backend/api/internal/repository/user"
)

// Reader interface
type Reader interface {
	List() ([]models.User, error)
	Get(email string) (models.User, error)
	GetFriendList(email string) (mod.FriendList, error)
	GetCommonFriends(email string, friend string) (mod.FriendList, error)
	GetRetrieveUpdates(sender string, mentions []*emailaddress.EmailAddress) (mod.RetrieveUpdates, error)
}

// Writer user writer
type Writer interface {
	CreateFriendship(email string, friend string) (mod.UserResponse, error)
	CreateSubscribe(requestor string, target string) (mod.UserResponse, error)
	CreateBlock(requestor string, target string) error
}

// UserInterface interface
type UserInterface interface {
	Reader
	Writer
}

// UserController: User Controller
type UserController struct {
	repo user.Repository
}

// NewUserController: Create new User Controller
func NewUserController(r user.Repository) *UserController {
	return &UserController{
		repo: r,
	}
}
