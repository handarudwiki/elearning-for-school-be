package models

import (
	"context"
	"time"

	"github.com/handarudwiki/models/dto"
)

type Event struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    uint      `json:"user_id" `
	User      *User     `json:"user"`
	Title     string    `json:"title"`
	Location  string    `json:"location"`
	Body      string    `json:"body"`
	Date      time.Time `json:"date"`
	Time      string    `json:"time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EventRepository interface {
	Create(ctx context.Context, event *Event) (*Event, error)
	FindByID(ctx context.Context, id int) (*Event, error)
	Update(ctx context.Context, event *Event, id int) (*Event, error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context, dto dto.QueryDTO) ([]*Event, int, error)
}
