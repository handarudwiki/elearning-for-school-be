package dto

type CreateAbcentDTO struct {
	UserID      int    `json:"-"`
	ScheduleID  int    `json:"schedule_id" validate:"required"`
	IsAbcent    bool   `json:"is_abcent"`
	Description string `json:"description"`
	Reason      int    `json:"reason"`
	Details     string `json:"details"`
}
