package services

import (
	"errors"
	"github.com/BetterToPractice/go-gin-setup/api/repositories"
	"github.com/BetterToPractice/go-gin-setup/lib"
	"github.com/BetterToPractice/go-gin-setup/models"
)

type UserService struct {
	db             lib.Database
	userRepository repositories.UserRepository
}

func NewUserService(db lib.Database, userRepository repositories.UserRepository) UserService {
	return UserService{
		db:             db,
		userRepository: userRepository,
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

func (c UserService) Query(params *models.UserQueryParams) (*models.UserPaginationResult, error) {
	return c.userRepository.Query(params)
}

func (c UserService) GetByUsername(username string) (*models.User, error) {
	return c.userRepository.GetByUsername(username)
}
