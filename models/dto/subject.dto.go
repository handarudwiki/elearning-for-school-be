package dto

type CreateSubjectDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type UpdateSubjectDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
