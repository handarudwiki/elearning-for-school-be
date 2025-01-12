package dto

type CreateUserDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Details  string `json:"details" validate:"required"`
}

type UpdateUserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Details  string `json:"details"`
	Password string `json:"password"`
	Role     int    `json:"role"`
	IsActive bool   `json:"is_active"`
}
