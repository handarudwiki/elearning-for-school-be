package dto

type InfoDto struct {
	UserID   uint   `json:"-"`
	Title    string `json:"title" validate:"required"`
	Body     string `json:"body" validate:"required"`
	Status   bool   `json:"status"`
	Settings string `json:"settings"`
}
