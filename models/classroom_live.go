package models

import "time"

type ClassroomLive struct {
	ID         uint      `json:"id"`
	ScheduleID uint      `json:"schedule_id"`
	Schedule   *Schedule `json:"schedule,omitempty" db:"-"`
	Body       string    `json:"body"`
	Settings   string    `json:"settings"`
	IsActive   bool      `json:"is_active"`
	Note       string    `json:"note"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
