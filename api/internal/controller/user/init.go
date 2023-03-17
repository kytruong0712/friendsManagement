package user

import (
	"github.com/mcnijman/go-emailaddress"

	"backend/api/internal/presenter"
	"backend/api/internal/repository/user"
	"backend/api/pkg/utils"
)

// Reader interface
type Reader interface {
	List() ([]*presenter.User, error)
	Get(email string) (*presenter.User, error)
	GetFriendList(email string) (*presenter.FriendList, error)
	GetCommonFriends(email string, friend string) (*presenter.FriendList, error)
	GetRetrieveUpdates(sender string, mentions []*emailaddress.EmailAddress) (*presenter.RetrieveUpdates, error)
}

// Writer user writer
type Writer interface {
	CreateFriendship(email string, friend string) (utils.JSONResponse, error)
	CreateSubscribe(requestor string, target string) (utils.JSONResponse, error)
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
