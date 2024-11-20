package repositories

import (
	"context"
	"math"

	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/models"
	"github.com/handarudwiki/models/dto"
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
		return nil, err
	}

	return info, err
}

func (r *infoRepository) Update(ctx context.Context, id int, info *models.Info) (*models.Info, error) {
	err := r.db.WithContext(ctx).Model(&models.Info{}).Where("id = ?", id).Updates(info).Error

	if err != nil {
		return &models.Info{}, err
	}

	return info, err
}

func (r *infoRepository) Delete(ctx context.Context, id int) error {
	err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Info{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *infoRepository) FindAll(ctx context.Context, dto dto.QueryDTO) ([]*models.Info, int, error) {
	var infos []*models.Info

	err := r.db.WithContext(ctx).
		Scopes(helpers.SearchTitle(*dto.Search)).
		Find(&infos).Error

	if err != nil {
		return nil, 0, err
	}

	totalPages := math.Ceil(float64(len(infos)) / float64(dto.Size))

	if err != nil {
		return nil, 0, err
	}

	return infos, int(totalPages), err
}
