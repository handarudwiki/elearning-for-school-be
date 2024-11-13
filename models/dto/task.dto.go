package dto

type CreateTaskDTO struct {
	Title    string `json:"title" validate:"required"`
	UserID   uint   `json:"user_id" validate:"required"`
	Type     string `json:"type" validate:"required"`
	Body     string `json:"body" validate:"required"`
	IsActive bool   `json:"is_active"`
	Settings string `json:"settings"`
	Deadline string `json:"deadline" validate:"required"`
}
