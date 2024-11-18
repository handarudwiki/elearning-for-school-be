package models

import "context"

type ClassroomSubject struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	ClassroomID uint       `json:"classroom_id"`
	Classroom   *Classroom `json:"classroom,omitempty" db:"-"`
	SubjectID   uint       `json:"subject_id"`
	Subject     *Subject   `json:"subject,omitempty" db:"-"`
	TeacherID   uint       `json:"teacher_id"`
	Teacher     *User      `json:"teacher,omitempty" db:"-"`
}

type ClassroomSubjectRepository interface {
	Create(ctx context.Context, classroomSubject *ClassroomSubject) (*ClassroomSubject, error)
	FindByTeacherID(ctx context.Context, teacherID int) ([]*ClassroomSubject, error)
}
