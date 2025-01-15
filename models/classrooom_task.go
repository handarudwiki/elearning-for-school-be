package models

import "context"

type ClassroomTask struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	ClassroomID uint       `json:"classroom_id"`
	Classroom   *Classroom `json:"classroom,omitempty" db:"-"`
	TaskID      uint       `json:"task_id"`
	Task        *Task      `json:"task,omitempty" db:"-"`
	TeacherID   uint       `json:"teacher_id"`
	Teacher     *User      `json:"teacher,omitempty" db:"-"`
	Body        string     `json:"body"`
}

type ClassroomTaskRepository interface {
	Create(ctx context.Context, classroomTask *ClassroomTask) (*ClassroomTask, error)
	FindByID(ctx context.Context, id int) (*ClassroomTask, error)
	FindByClassroomID(ctx context.Context, classroomID int) ([]*ClassroomTask, error)
}
