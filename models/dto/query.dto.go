package dto

type QueryDTO struct {
	Search    *string `json:"search"`
	Page      int     `json:"page"`
	Size      int     `json:"size"`
	Is_active *bool   `json:"is_active"`
	Role      *int    `json:"role"`
}
