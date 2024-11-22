package dto

type StandartDTO struct {
	SubjectID  uint   `json:"subject_id" validate:"required,gt=0"`
	TeacherID  uint   `json:"-"`
	Type       string `json:"type" validate:"required"`
	Body       string `json:"body" validate:"required"`
	Code       string `json:"code" validate:"required"`
	StandartID uint   `json:"standart_id" `
}
