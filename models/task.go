package models

type Task struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	UserID    uint   `json:"user_id"`
	User      User   `json:"user"`
	Type      int    `json:"type"`
	Body      string `json:"body"`
	IsActive  bool   `json:"is_active"`
	Settings  string `json:"settings"`
	Deadline  string `json:"deadline"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
