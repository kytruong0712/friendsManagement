package user

import (
	"fmt"

	"backend/api/pkg/constants"
	"backend/api/pkg/utils"
)

// CreateFriendship: create a friend connection between two email addresses.
func (c UserController) CreateFriendship(email string, friend string) (utils.JSONResponse, error) {
	errorResp := utils.JSONResponse{
		Success: false,
		Message: fmt.Sprintf("Error while creating friendship between %s and %s", email, friend),
	}

	err := c.repo.CreateRelationship(email, friend, constants.AddFriendToExistingFriendsArray)
	if err != nil {
		return errorResp, err
	}

	err = c.repo.CreateRelationship(email, friend, constants.AddFriendToNullFriendsArray)
	if err != nil {
		return errorResp, err
	}

	err = c.repo.CreateRelationship(friend, email, constants.AddFriendToExistingFriendsArray)
	if err != nil {
		return errorResp, err
	}

	err = c.repo.CreateRelationship(friend, email, constants.AddFriendToNullFriendsArray)
	if err != nil {
		return errorResp, err
	}

	data, er := c.repo.Get(email)
	if er != nil {
		return errorResp, er
	}

	if len(data.Blocks) == 0 {
		return utils.JSONResponse{
			Success: true,
			Message: "create a friend connection successfully",
		}, nil
	} else {
		isBlocked, e := c.repo.IsBlock(email, friend)
		if e != nil {
			return errorResp, er
		}

		if isBlocked.Blocked {
			resp := utils.JSONResponse{
				Success: false,
				Message: fmt.Sprintf("Cannot add friend because %s has blocked %s", email, friend),
			}
			return resp, nil
		} else {
			resp := utils.JSONResponse{
				Success: true,
				Message: "create a friend connection successfully",
			}
			return resp, nil
		}
	}
}

// CreateSubscribe: subscribe to updates from an email address.
func (c UserController) CreateSubscribe(requestor string, target string) (utils.JSONResponse, error) {
	errorResp := utils.JSONResponse{
		Success: false,
		Message: fmt.Sprintf("Error while creating Subscribe between %s has blocked %s", requestor, target),
	}

	err := c.repo.CreateRelationship(requestor, target, constants.AddSubscribeToExistingSubscribeArray)
	if err != nil {
		return errorResp, err
	}

	err = c.repo.CreateRelationship(requestor, target, constants.AddSubscribeToNullSubscribeArray)
	if err != nil {
		return errorResp, err
	}

	data, er := c.repo.Get(requestor)
	if er != nil {
		return errorResp, er
	}

	if len(data.Blocks) == 0 {
		return utils.JSONResponse{
			Success: true,
			Message: "create a subscribe successfully",
		}, nil
	} else {
		isBlocked, e := c.repo.IsBlock(requestor, target)
		if e != nil {
			return errorResp, er
		}

		if isBlocked.Blocked {
			resp := utils.JSONResponse{
				Success: false,
				Message: fmt.Sprintf("Cannot subscribe because %s has blocked %s", requestor, target),
			}
			return resp, nil
		} else {
			resp := utils.JSONResponse{
				Success: true,
				Message: "create a subscribe successfully",
			}
			return resp, nil
		}
	}
}

// CreateBlock: block updates from an email address.
func (c UserController) CreateBlock(requestor string, target string) error {
	err := c.repo.CreateRelationship(requestor, target, constants.AddSubscribeToNullSubscribeArray)
	if err != nil {
		return err
	}

	err = c.repo.CreateRelationship(requestor, target, constants.AddBlockToNullSubscribeArray)
	if err != nil {
		return err
	}

	return nil
}
