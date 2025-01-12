package models

import (
	"context"
	"time"

	"github.com/handarudwiki/models/dto"
)

type LectureComment struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	LectureID uint      `json:"lecture_id"`
	Content   string    `json:"content"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LectureCommentRepository interface {
	Create(ctx context.Context, comment *LectureComment) (*LectureComment, error)
	FindByLectureID(ctx context.Context, lectureID int, query dto.QueryDTO) ([]*LectureComment, int64, error)
}
