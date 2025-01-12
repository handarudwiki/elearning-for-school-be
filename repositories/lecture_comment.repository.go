package repositories

import (
	"context"

	"github.com/handarudwiki/models"
	"gorm.io/gorm"
)

type LectureCommentRepository struct {
	db *gorm.DB
}

func NewLectureComment(db *gorm.DB) models.LectureCommentRepository {
	return &LectureCommentRepository{
		db: db,
	}
}

func (r *LectureCommentRepository) Create(ctx context.Context, comment *models.LectureComment) (*models.LectureComment, error) {
	err := r.db.Create(comment).Error
	if err != nil {
		return nil, err
	}

	return comment, nil
}
