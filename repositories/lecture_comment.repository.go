package repositories

import (
	"context"

	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/models"
	"github.com/handarudwiki/models/dto"
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

func (r *LectureCommentRepository) FindByLectureID(ctx context.Context, lectureID int, query dto.QueryDTO) (comments []*models.LectureComment, count int64, err error) {
	err = r.db.Debug().Where("lecture_id = ?", lectureID).Scopes(helpers.Paginate(query.Page, query.Size)).
		Order("created_at DESC").Find(&comments).Error
	if err != nil {
		return
	}

	err = r.db.Model(&models.LectureComment{}).Where("lecture_id = ?", lectureID).Count(&count).Error

	if err != nil {
		return nil, 0, err
	}

	return comments, count, nil
}
