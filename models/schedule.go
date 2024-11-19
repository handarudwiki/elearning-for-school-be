package models

import (
	"context"
	"time"
)

type Schedule struct {
	ID                 uint              `json:"id"`
	ClassroomSubjectID uint              `json:"classroom_subject_id"`
	ClassroomSubject   *ClassroomSubject `json:"classroom_subject,omitempty" db:"-"`
	Day                int               `json:"day"`
	StartTime          string            `json:"start_at"`
	EndTime            string            `json:"end_at"`
	CreatedAt          time.Time         `json:"created_at"`
	UpdatedAt          time.Time         `json:"updated_at"`
}

type ScheduleRepository interface {
	Create(ctc context.Context, schedule *Schedule) (*Schedule, error)
	FindByID(ctx context.Context, id int) (*Schedule, error)
	FindByClassroomSubjectID(ctx context.Context, classroomSubjectID int) ([]*Schedule, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, schedule *Schedule) (*Schedule, error)
	GetScheduleByday(ctx context.Context, day, teacherID int) ([]*Schedule, error)
}
