package response

import "github.com/handarudwiki/models"

type AbcentResponse struct {
	ID          uint          `json:"id"`
	UserID      uint          `json:"user_id"`
	ScheduleID  uint          `json:"schedule_id"`
	IsAbcent    bool          `json:"is_abcent"`
	Description string        `json:"description"`
	User        *UserResponse `json:"user"`
	Details     string        `json:"details"`
	Reason      int           `json:"reason"`
}

func ToAbcentResponse(abcent *models.Abcent) *AbcentResponse {

	var user UserResponse

	if abcent.User != nil {
		user = ToUserResponse(abcent.User)
	}

	return &AbcentResponse{
		ID:          abcent.ID,
		UserID:      abcent.UserID,
		ScheduleID:  abcent.ScheduleID,
		IsAbcent:    abcent.IsAbcent,
		Description: abcent.Description,
		Details:     abcent.Details,
		Reason:      abcent.Reason,
		User:        &user,
	}
}

func ToAbcentResponses(abcents []*models.Abcent) []*AbcentResponse {
	var responses []*AbcentResponse
	for _, abcent := range abcents {
		responses = append(responses, ToAbcentResponse(abcent))
	}
	return responses
}
