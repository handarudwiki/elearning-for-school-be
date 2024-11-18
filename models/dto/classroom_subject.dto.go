package dto

type CreateClassrooomSubject struct {
	ClassroomID uint `json:"classroom_id" validate:"required,gt=0"`
	SubjectID   uint `json:"subject_id" validate:"required,gt=0"`
	TeacherID   uint `json:"-"`
}
