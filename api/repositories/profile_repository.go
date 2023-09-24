package repositories

import (
	"github.com/BetterToPractice/go-gin-setup/lib"
	"github.com/BetterToPractice/go-gin-setup/models"
)

type ProfileRepository struct {
	db lib.Database
}

func NewProfileRepository(db lib.Database) ProfileRepository {
	return ProfileRepository{
		db: db,
	}
}

func (r ProfileRepository) DeleteByUserID(userID string) error {
	profile := new(models.Profile)
	if err := r.db.ORM.Model(profile).Where("user_id = ?", userID).Delete(profile).Error; err != nil {
		return err
	}
	return nil
}
