package controller

import (
	"github.com/mcnijman/go-emailaddress"

	"backend/api/internal/presenter"
	"backend/api/pkg/utils"
)

// List: Get all users
func (c UserController) List() ([]*presenter.User, error) {
	data, err := c.repo.List()

	if err != nil {
		return nil, err
	}

	var users []*presenter.User
	for _, d := range data {
		users = append(users, &presenter.User{
			ID:        d.ID,
			Name:      d.Name,
			Email:     d.Email,
			Friends:   d.Friends,
			Subscribe: d.Subscribe,
			Blocks:    d.Blocks,
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		})
	}

	return users, nil
}

// Get: Get single user by email
func (c UserController) Get(email string) (*presenter.User, error) {
	data, err := c.repo.Get(email)

	if err != nil {
		return nil, err
	}

	user := &presenter.User{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		Friends:   data.Friends,
		Subscribe: data.Subscribe,
		Blocks:    data.Blocks,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}

	return user, nil
}

// GetFriendList: retrieve the friends list for an email address.
func (c UserController) GetFriendList(email string) (*presenter.FriendList, error) {
	data, err := c.repo.Get(email)

	if err != nil {
		return nil, err
	}

	count := len(data.Friends)

	friendsList := make([]string, 0)
	if count > 0 {
		friendsList = data.Friends
	}

	resp := &presenter.FriendList{
		Success: true,
		Friends: friendsList,
		Count:   count,
	}

	return resp, nil
}

// GetCommonFriends: retrieve the common friends list between two email addresses.
func (c UserController) GetCommonFriends(email string, friend string) (*presenter.FriendList, error) {
	users1, err1 := c.repo.Get(email)
	if err1 != nil {
		return nil, err1
	}

	users2, err2 := c.repo.Get(friend)
	if err2 != nil {
		return nil, err2
	}

	friends1 := make([]string, 0)
	if len(users1.Friends) > 0 {
		friends1 = users1.Friends
	}

	friends2 := make([]string, 0)
	if len(users2.Friends) > 0 {
		friends2 = users2.Friends
	}

	temp_intersect := utils.HashGeneric(friends1, friends2)
	intersect := make([]string, 0)
	for _, value := range temp_intersect {
		if value != email && value != friend {
			intersect = append(intersect, value)
		}
	}

	resp := &presenter.FriendList{
		Success: true,
		Friends: intersect,
		Count:   len(intersect),
	}

	return resp, nil
}

// GetRetrieveUpdates: retrieve all email addresses that can receive updates from an email address.
func (c UserController) GetRetrieveUpdates(sender string, mentions []*emailaddress.EmailAddress) (*presenter.RetrieveUpdates, error) {
	data, err := c.repo.Get(sender)
	if err != nil {
		return nil, err
	}

	retrieveList := make([]string, 0)
	retrieveList = utils.AppendWithoutDuplicate(retrieveList, data.Friends)
	retrieveList = utils.AppendWithoutDuplicate(retrieveList, data.Subscribe)
	for _, m := range mentions {
		retrieveList = utils.AppendWithoutDuplicate(retrieveList, []string{m.LocalPart + "@" + m.Domain})
	}

	retrieveList = utils.FindMissing(retrieveList, data.Blocks)

	resp := &presenter.RetrieveUpdates{
		Success:    true,
		Message:    "retrieve updates successfully",
		Recipients: retrieveList,
	}

	return resp, nil
}
