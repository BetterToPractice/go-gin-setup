package services

import (
	"errors"
	"github.com/BetterToPractice/go-gin-setup/api/dto"
	"github.com/BetterToPractice/go-gin-setup/api/mails"
	"github.com/BetterToPractice/go-gin-setup/api/repositories"
	appErrors "github.com/BetterToPractice/go-gin-setup/errors"
	"github.com/BetterToPractice/go-gin-setup/lib"
	"github.com/BetterToPractice/go-gin-setup/models"
	"strconv"
)

type UserService struct {
	db                lib.Database
	userRepository    repositories.UserRepository
	profileRepository repositories.ProfileRepository
	authMail          mails.AuthMail
}

func NewUserService(db lib.Database, userRepository repositories.UserRepository, profileRepository repositories.ProfileRepository, authMail mails.AuthMail) UserService {
	return UserService{
		db:                db,
		userRepository:    userRepository,
		profileRepository: profileRepository,
		authMail:          authMail,
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

func (s UserService) Query(params *models.UserQueryParams) (*models.UserPaginationResult, error) {
	return s.userRepository.Query(params)
}

func (s UserService) GetByUsername(username string) (*models.User, error) {
	return s.userRepository.GetByUsername(username)
}

func (s UserService) Register(params *dto.RegisterRequest) (*models.User, error) {
	user := &models.User{
		Username: params.Username,
		Password: params.Password,
		Email:    params.Email,
	}

	if err := s.userRepository.Create(user); err != nil {
		return nil, errors.Join(appErrors.DatabaseInternalError, err)
	}
	return user, nil
}

func (s UserService) Delete(user *models.User) error {
	if err := s.profileRepository.DeleteByUserID(strconv.Itoa(int(user.ID))); err != nil {
		return err
	}
	if err := s.userRepository.Delete(user); err != nil {
		return err
	}

	return nil
}
