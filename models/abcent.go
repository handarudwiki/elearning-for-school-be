package models

import (
	"context"
	"time"
)

type Abcent struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	UserID      uint      `json:"user_id"`
	ScheduleID  uint      `json:"schedule_id"`
	IsAbcent    bool      `json:"is_abcent"`
	Reason      int       `json:"reason"`
	Description string    `json:"description"`
	Details     string    `json:"details"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AbcentRepository interface {
	Create(ctx context.Context, abcent *Abcent) (*Abcent, error)
}
