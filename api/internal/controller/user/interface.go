package controller

import (
	"github.com/mcnijman/go-emailaddress"

	"backend/api/internal/models"
	"backend/api/internal/presenter"
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

// Repository interface
type Repository interface {
	List() ([]models.User, error)
	Get(email string) (models.User, error)
	CreateRelationship(email string, friend string, stmt string) error
	IsBlock(requestor string, target string) (*presenter.IsBlock, error)
}
