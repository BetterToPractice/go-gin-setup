package services

import (
	"errors"
	"github.com/BetterToPractice/go-gin-setup/lib"
	"github.com/BetterToPractice/go-gin-setup/models"
)

type UserService struct {
	db lib.Database
}

func NewUserService(db lib.Database) UserService {
	return UserService{
		db: db,
	}
}

func (c UserService) Verify(username, password string) (*models.User, error) {
	user := new(models.User)
	err := c.db.ORM.First(&user, "username = ?", username).Error
	if err != nil || user.Password != password {
		return nil, errors.New("username or password not match")
	}
	return user, err
}
