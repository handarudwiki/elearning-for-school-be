package repositories

import (
	"context"

	"github.com/handarudwiki/models"
	"gorm.io/gorm"
)

type infoRepository struct {
	db *gorm.DB
}

func NewInfo(db *gorm.DB) models.InfoRepository {
	return &infoRepository{
		db: db,
	}
}

func (r *infoRepository) Create(ctx context.Context, info *models.Info) (*models.Info, error) {
	err := r.db.WithContext(ctx).Create(info).Error

	if err != nil {
		return &models.Info{}, err
	}

	return info, err
}

func (r *infoRepository) FindByID(ctx context.Context, id int) (*models.Info, error) {
	var info *models.Info

	err := r.db.WithContext(ctx).Where("id = ?", id).Preload("User").First(info).Error

	if err != nil {
		return &models.Info{}, err
	}

	return info, err
}
