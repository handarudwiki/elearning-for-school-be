package dto

import "time"

type EventDto struct {
	Title    string    `json:"title" validate:"required"`
	UserID   uint      `json:"-"`
	Location string    `json:"location" validate:"required"`
	Body     string    `json:"body" validate:"required"`
	Date     time.Time `json:"date" validate:"required,future_date"`
	Time     string    `json:"time" validate:"required,time"`
}
