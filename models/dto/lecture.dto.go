package dto

type CreateLectureDTO struct {
	SubjectID uint   `json:"subject_id" validate:"required,gt=0"`
	Title     string `json:"title" validate:"required"`
	Body      string `json:"body" validate:"required"`
	Addition  string `json:"additions" validate:"required"`
	UserID    uint   `json:"user_id,omitempty"`
	Settings  string `json:"settings" validate:"required"`
}
type UpdateLectureDTO struct {
	SubjectID uint   `json:"subject_id" validate:"required,gt=0"`
	Title     string `json:"title" validate:"required"`
	Body      string `json:"body" validate:"required"`
	Addition  string `json:"additions" validate:"required"`
	UserID    uint   `json:"user_id,omitempty"`
	Settings  string `json:"settings" validate:"required"`
}
