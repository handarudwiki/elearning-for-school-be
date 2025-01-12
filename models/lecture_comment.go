package models

import (
	"context"
	"time"
)

type LectureComment struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	LectureID uint      `json:"lecture_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LectureCommentRepository interface {
	Create(ctx context.Context, comment *LectureComment) (*LectureComment, error)
}
