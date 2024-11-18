package models

import (
	"context"
	"time"

	"github.com/handarudwiki/models/dto"
)

type Task struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	UserID    uint      `json:"user_id"`
	User      *User     `json:"user,omitempty" db:"-"`
	Type      int       `json:"type"`
	Body      string    `json:"body"`
	IsActive  bool      `json:"is_active"`
	Settings  string    `json:"settings"`
	Deadline  time.Time `json:"deadline"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskRepository interface {
	Create(ctx context.Context, task *Task) (*Task, error)
	FindByID(ctx context.Context, id int) (*Task, error)
	Update(ctx context.Context, id int, task *Task) (*Task, error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context, query dto.QueryDTO) ([]*Task, int, error)
}
