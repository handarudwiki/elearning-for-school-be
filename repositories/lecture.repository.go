package repositories

import (
	"context"

	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/models"
	"github.com/handarudwiki/models/dto"
	"gorm.io/gorm"
)

type lectureRepository struct {
	db *gorm.DB
}

func NewLecture(db *gorm.DB) models.LectureRepository {
	return &lectureRepository{
		db: db,
	}
}

func (r *lectureRepository) FindByID(ctx context.Context, id int) (res *models.Lecture, err error) {
	var lecture models.Lecture
	err = r.db.Where("id = ?", id).Preload("User").Preload("Subject").First(&lecture).Error
	if err != nil {
		return nil, err
	}

	return &lecture, nil
}

func (r *lectureRepository) Create(ctx context.Context, lecture *models.Lecture) (*models.Lecture, error) {
	err := r.db.Create(lecture).Error
	if err != nil {
		return nil, err
	}

	return lecture, nil
}

func (r *lectureRepository) Update(ctx context.Context, lecture *models.Lecture, id int) (*models.Lecture, error) {
	err := r.db.Model(&models.Lecture{}).Where("id = ?", id).Updates(lecture).Error
	if err != nil {
		return nil, err
	}

	return lecture, nil
}

func (r *lectureRepository) Delete(ctx context.Context, id int) error {
	err := r.db.Where("id = ?", id).Delete(&models.Lecture{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *lectureRepository) FindAll(ctx context.Context, dto dto.QueryDTO) ([]*models.Lecture, int, error) {
	var lectures []*models.Lecture

	var totalPage int64

	err := r.db.Model(&models.Lecture{}).Scopes(helpers.FilterIsActive(dto.Is_active), helpers.SearchTitle(*dto.Search)).Count(&totalPage).Error

	if err != nil {
		return nil, 0, err
	}

	err = r.db.Scopes(helpers.Paginate(dto.Page, dto.Size), helpers.FilterIsActive(dto.Is_active), helpers.SearchTitle(*dto.Search)).Preload("User").Preload("Subject").Find(&lectures).Error

	if err != nil {
		return nil, 0, err
	}

	return lectures, int(totalPage), nil
}
