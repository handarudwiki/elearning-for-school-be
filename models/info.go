package models

import (
	"context"
	"time"

	"github.com/handarudwiki/models/dto"
)

type Info struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	UserID    int       `json:"user_id"`
	User      *User     `json:"user,omitempty" db:"-"`
	Body      string    `json:"body"`
	Status    bool      `json:"status"`
	Settings  string    `json:"settings"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InfoRepository interface {
	Create(ctx context.Context, info *Info) (*Info, error)
	FindByID(ctx context.Context, id int) (*Info, error)
	Update(ctx context.Context, id int, info *Info) (*Info, error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context, dto dto.QueryDTO) ([]*Info, int, error)
}
