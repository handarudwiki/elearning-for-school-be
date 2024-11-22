package models

import (
	"context"
	"time"
)

type Standart struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	TeacherID  uint      `json:"teacher_id"`
	Teacher    *User     `json:"teacher,omitempty" db:"-"`
	StandartID uint      `json:"standart_id"`
	SubjectID  uint      `json:"subject_id"`
	Subject    *User     `json:"subject,omitempty" db:"-"`
	Type       string    `json:"type"`
	Body       string    `json:"body"`
	Code       string    `json:"code"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type StandartRepository interface {
	Create(ctx context.Context, standart *Standart) (*Standart, error)
	FindByID(ctx context.Context, id int) (*Standart, error)
	Update(ctx context.Context, standart *Standart, id int) (*Standart, error)
	Delete(ctx context.Context, id int) error
}
