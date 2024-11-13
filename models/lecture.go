package models

import (
	"context"
	"time"
)

type Lecture struct {
	ID        uint      `json:"id"`
	SubjectID uint      `json:"subject_id"`
	Subject   Subject   `json:"subject"`
	UserID    uint      `json:"user_id"`
	User      User      `json:"user"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	IsActive  bool      `json:"is_active"`
	Addition  string    `json:"additions"`
	Settings  string    `json:"settings"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LectureRepository interface {
	FindByID(ctx context.Context, id int) (*Lecture, error)
	Create(ctx context.Context, lecture *Lecture) (*Lecture, error)
}
