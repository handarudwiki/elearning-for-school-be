package models

import (
	"context"
	"time"

	"github.com/handarudwiki/models/dto"
)

type Abcent struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	UserID      uint      `json:"user_id"`
	User        *User     `gorm:"foreignKey:UserID" json:"user"`
	ScheduleID  uint      `json:"schedule_id"`
	IsAbcent    bool      `json:"is_abcent"`
	Reason      int       `json:"reason"`
	Description string    `json:"description"`
	Details     string    `json:"details"`
	CreatedAt   time.Time `json:"created_a t"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AbcentRepository interface {
	Create(ctx context.Context, abcent *Abcent) (*Abcent, error)
	FindByScheduleIDToday(ctx context.Context, scheduleID int, date string, query dto.QueryDTO) (abcents []*Abcent, count int64, err error)
}
