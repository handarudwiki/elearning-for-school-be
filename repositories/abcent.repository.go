package repositories

import (
	"context"

	"github.com/handarudwiki/models"
	"gorm.io/gorm"
)

type abcentRepository struct {
	db *gorm.DB
}

func NewAbcent(db *gorm.DB) models.AbcentRepository {
	return &abcentRepository{
		db: db,
	}
}

func (r *abcentRepository) Create(ctx context.Context, abcent *models.Abcent) (*models.Abcent, error) {
	err := r.db.Create(abcent).Error
	if err != nil {
		return nil, err
	}

	return abcent, nil
}
