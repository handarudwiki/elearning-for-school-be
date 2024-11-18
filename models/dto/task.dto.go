package dto

import "time"

type CreateTaskDTO struct {
	Title    string    `json:"title" validate:"required"`
	UserID   uint      `json:"-"`
	Type     int       `json:"type" validate:"required,gt=0"`
	Body     string    `json:"body" validate:"required"`
	IsActive bool      `json:"is_active"`
	Settings string    `json:"settings"`
	Deadline time.Time `json:"deadline" validate:"required"`
}

type UpdateTaskDTO struct {
	Title    string    `json:"title" validate:"required"`
	UserID   uint      `json:"-"`
	Type     int       `json:"type" validate:"required,gt=0"`
	Body     string    `json:"body" validate:"required"`
	IsActive bool      `json:"is_active"`
	Settings string    `json:"settings"`
	Deadline time.Time `json:"deadline" validate:"required,future_date"`
}
