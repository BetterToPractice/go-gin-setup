package policies

import (
	appErrors "github.com/BetterToPractice/go-gin-setup/errors"
	"github.com/BetterToPractice/go-gin-setup/models"
)

type UserPolicy struct {
}

func NewUserPolicy() UserPolicy {
	return UserPolicy{}
}

func (p UserPolicy) CanViewList(_ *models.User) (bool, error) {
	return true, nil
}

func (p UserPolicy) CanViewDetail(_ *models.User, _ *models.Post) (bool, error) {
	return true, nil
}

func (p UserPolicy) CanCreate(user *models.User) (bool, error) {
	if user == nil {
		return false, appErrors.Unauthorized
	}
	return true, nil
}

func (p UserPolicy) CanUpdate(user *models.User, post *models.Post) (bool, error) {
	if user == nil {
		return false, appErrors.Unauthorized
	}
	if post.UserID != user.ID {
		return false, appErrors.Forbidden
	}
	return true, nil
}

func (p UserPolicy) CanDelete(loggedInUser *models.User, user *models.User) (bool, error) {
	if loggedInUser == nil {
		return false, appErrors.Unauthorized
	}
	if loggedInUser.ID != user.ID {
		return false, appErrors.Forbidden
	}
	return true, nil
}
