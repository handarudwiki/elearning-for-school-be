package models

import (
	"context"
	"time"
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
}
