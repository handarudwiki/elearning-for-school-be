package repositories

import (
	"context"
	"fmt"
	"math"

	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/models"
	"github.com/handarudwiki/models/dto"
	"gorm.io/gorm"
)

type standartRepository struct {
	db *gorm.DB
}

func NewStandart(db *gorm.DB) models.StandartRepository {
	return &standartRepository{
		db: db,
	}
}

func (r *standartRepository) FindByID(ctx context.Context, id int) (res *models.Standart, err error) {
	var standart models.Standart
	err = r.db.Where("id = ?", id).Preload("Teacher").First(&standart).Error
	if err != nil {
		return nil, err
	}

	return &standart, nil
}

func (r *standartRepository) Create(ctx context.Context, standart *models.Standart) (*models.Standart, error) {

	fmt.Println(standart.TeacherID)

	err := r.db.Create(standart).Error
	if err != nil {
		return nil, err
	}

	return standart, nil
}

func (r *standartRepository) Update(ctx context.Context, standart *models.Standart, id int) (*models.Standart, error) {
	err := r.db.Model(&standart).Where("id = ?", id).Updates(standart).Error
	if err != nil {
		return nil, err
	}

	return standart, nil
}

func (r *standartRepository) Delete(ctx context.Context, id int) error {
	err := r.db.Where("id = ?", id).Delete(&models.Standart{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *standartRepository) FindAll(ctx context.Context, dto *dto.QueryDTO) ([]*models.Standart, int, error) {
	var standarts []*models.Standart
	var count int64

	err := r.db.Scopes(helpers.Paginate(dto.Page, dto.Size)).Preload("Teacher").Preload("Subject").Find(&standarts).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.db.Model(&models.Standart{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	totalPage := math.Ceil(float64(count) / float64(dto.Size))

	return standarts, int(totalPage), nil
}
