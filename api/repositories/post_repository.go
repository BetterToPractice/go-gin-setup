package repositories

import (
	"errors"
	dto2 "github.com/BetterToPractice/go-gin-setup/api/dto"
	appErrors "github.com/BetterToPractice/go-gin-setup/errors"
	"github.com/BetterToPractice/go-gin-setup/lib"
	"github.com/BetterToPractice/go-gin-setup/models"
)

type PostRepository struct {
	db lib.Database
}

func NewPostRepository(db lib.Database) PostRepository {
	return PostRepository{
		db: db,
	}
}

func (r PostRepository) Query(params *dto2.PostQueryParam) (*dto2.PostPaginationResponse, error) {
	db := r.db.ORM.Preload("User").Model(&models.Posts{})

	var list models.Posts
	pagination, err := QueryPagination(db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.Join(appErrors.DatabaseInternalError, err)
	}

	qr := &dto2.PostPaginationResponse{
		Pagination: pagination,
	}
	qr.Serializer(&list)

	return qr, nil
}

func (r PostRepository) Get(id string) (*models.Post, error) {
	post := new(models.Post)
	if ok, err := QueryOne(r.db.ORM.Preload("User").Model(post).Where("id = ?", id), post); err != nil {
		return nil, err
	} else if !ok {
		return nil, appErrors.DatabaseRecordNotFound
	}
	return post, nil
}

func (r PostRepository) Create(post *models.Post) error {
	if err := r.db.ORM.Model(post).Create(post).Error; err != nil {
		return errors.Join(appErrors.DatabaseInternalError, err)
	}
	return nil
}

func (r PostRepository) Update(post *models.Post) error {
	if err := r.db.ORM.Model(post).Updates(post).Error; err != nil {
		return errors.Join(appErrors.DatabaseInternalError, err)
	}
	return nil
}

func (r PostRepository) Delete(post *models.Post) error {
	if err := r.db.ORM.Model(post).Where("id = ?", post.ID).Delete(post).Error; err != nil {
		return errors.Join(appErrors.DatabaseInternalError, err)
	}
	return nil
}
