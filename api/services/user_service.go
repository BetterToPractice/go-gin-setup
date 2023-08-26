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
	if err != nil || user.Password != models.HashPassword(password) {
		return nil, errors.New("username or password not match")
	}
	return user, err
}

func (c UserService) Register(username, password, email string) (*models.User, error) {
	user := &models.User{
		Username: username,
		Password: models.HashPassword(password),
		Email:    email,
	}
	err := c.db.ORM.Create(&user).Error
	return user, err
}

func (c UserService) Query() (*[]models.User, error) {
	var users = new([]models.User)
	err := c.db.ORM.Find(&users).Error
	return users, err
}

func (c UserService) GetByUsername(username string) (*models.User, error) {
	user := new(models.User)
	err := c.db.ORM.First(&user, "username = ?", username).Error
	return user, err
}
