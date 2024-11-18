package models

import (
	"context"
	"time"

	"github.com/handarudwiki/models/dto"
)

type Subject struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Settings    string    `json:"settings"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SubjectRepository interface {
	FindByID(ctx context.Context, id int) (subject *Subject, err error)
	Create(subject *Subject) (*Subject, error)
	FindAll(ctx context.Context, dto dto.QueryDTO) (subjects []*Subject, totalPage int64, err error)
	Update(ctx context.Context, id int, subject *Subject) (*Subject, error)
	Delete(ctx context.Context, id int) error
}
