package repositories

import (
	"context"

	"github.com/handarudwiki/models"
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
