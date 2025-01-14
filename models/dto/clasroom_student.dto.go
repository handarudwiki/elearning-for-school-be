package dto

type CreateClassroomStudentDTO struct {
	ClassroomID int `json:"classroom_id" validate:"required"`
	StudentID   int `json:"student_id" validate:"required"`
}
