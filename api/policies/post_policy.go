package policies

import (
	appErrors "github.com/BetterToPractice/go-gin-setup/errors"
	"github.com/BetterToPractice/go-gin-setup/models"
)

type PostPolicy struct {
}

func NewPostPolicy() PostPolicy {
	return PostPolicy{}
}

func (p PostPolicy) CanViewList(_ *models.User) (bool, error) {
	return true, nil
}

func (p PostPolicy) CanViewDetail(_ *models.User, _ *models.Post) (bool, error) {
	return true, nil
}

func (p PostPolicy) CanCreate(user *models.User) (bool, error) {
	if user == nil {
		return false, appErrors.Unauthorized
	}
	return true, nil
}

func (p PostPolicy) CanUpdate(user *models.User, post *models.Post) (bool, error) {
	if user == nil {
		return false, appErrors.Unauthorized
	}
	if post.UserID != user.ID {
		return false, appErrors.Forbidden
	}
	return true, nil
}

func (p PostPolicy) CanDelete(user *models.User, post *models.Post) (bool, error) {
	if user == nil {
		return false, appErrors.Unauthorized
	}
	if post.UserID != user.ID {
		return false, appErrors.Forbidden
	}
	return true, nil
}
