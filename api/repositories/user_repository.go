package repositories

import (
	"errors"
	appErrors "github.com/BetterToPractice/go-gin-setup/errors"
	"github.com/BetterToPractice/go-gin-setup/lib"
	"github.com/BetterToPractice/go-gin-setup/models"
)

type UserRepository struct {
	db lib.Database
}

func NewUserRepository(db lib.Database) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) Query(params *models.UserQueryParams) (*models.UserPaginationResult, error) {
	db := r.db.ORM.Preload("Profile").Model(&models.User{})
	list := make(models.Users, 0)

	pagination, err := QueryPagination(db, params.PaginationParam, &list)
	if err != nil {
		return nil, err
	}

	qr := &models.UserPaginationResult{
		Pagination: pagination,
		List:       list,
	}

	return qr, nil
}

func (r UserRepository) GetByUsername(username string) (*models.User, error) {
	user := new(models.User)

	if ok, err := QueryOne(r.db.ORM.Preload("Profile").Model(user).Where("username = ?", username), user); err != nil {
		return nil, appErrors.RecordNotFound
	} else if !ok {
		return nil, errors.New("not Found")
	}

	return user, nil
}

func (r UserRepository) Delete(username string) error {
	user := new(models.User)
	if err := r.db.ORM.Model(user).Where("username = ?", username).Delete(user).Error; err != nil {
		return err
	}
	return nil
}
