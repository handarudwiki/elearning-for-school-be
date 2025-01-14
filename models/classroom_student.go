package models

import (
	"context"
	"time"
)

type ClassroomStudent struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	StudentID   uint       `json:"student_id"`
	ClassroomID uint       `json:"classroom_id"`
	Student     *User      `gorm:"foreignKey:StudentID" json:"student"`
	Classroom   *Classroom `gorm:"foreignKey:ClassroomID" json:"classroom"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type ClassroomStudentRepository interface {
	FindByClassroomID(ctx context.Context, classroomID int) ([]*ClassroomStudent, error)
	Create(ctx context.Context, classroomStudent *ClassroomStudent) (*ClassroomStudent, error)
}
