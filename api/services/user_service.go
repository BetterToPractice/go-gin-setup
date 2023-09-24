package services

import (
	"errors"
	"github.com/BetterToPractice/go-gin-setup/api/repositories"
	"github.com/BetterToPractice/go-gin-setup/lib"
	"github.com/BetterToPractice/go-gin-setup/models"
	"strconv"
)

type UserService struct {
	db                lib.Database
	userRepository    repositories.UserRepository
	profileRepository repositories.ProfileRepository
}

func NewUserService(db lib.Database, userRepository repositories.UserRepository, profileRepository repositories.ProfileRepository) UserService {
	return UserService{
		db:                db,
		userRepository:    userRepository,
		profileRepository: profileRepository,
	}
}

func (s UserService) Verify(username, password string) (*models.User, error) {
	user := new(models.User)
	err := s.db.ORM.First(&user, "username = ?", username).Error
	if err != nil || user.Password != models.HashPassword(password) {
		return nil, errors.New("username or password not match")
	}
	return user, err
}

func (s UserService) Register(username, password, email string) (*models.User, error) {
	user := &models.User{
		Username: username,
		Password: models.HashPassword(password),
		Email:    email,
	}
	err := s.db.ORM.Create(&user).Error
	return user, err
}

func (s UserService) Query(params *models.UserQueryParams) (*models.UserPaginationResult, error) {
	return s.userRepository.Query(params)
}

func (s UserService) GetByUsername(username string) (*models.User, error) {
	return s.userRepository.GetByUsername(username)
}

func (s UserService) Delete(username string) error {
	user, err := s.userRepository.GetByUsername(username)
	if err != nil {
		return err
	}

	if err := s.profileRepository.DeleteByUserID(strconv.Itoa(int(user.ID))); err != nil {
		return err
	}
	if err := s.userRepository.Delete(username); err != nil {
		return err
	}

	return nil
}
