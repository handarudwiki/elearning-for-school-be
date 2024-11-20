package response

import "github.com/handarudwiki/models"

type InfoResponse struct {
	ID       int           `json:"id"`
	UserID   int           `json:"user_id"`
	User     *UserResponse `json:"user"`
	Title    string        `json:"title"`
	Body     string        `json:"body"`
	Status   bool          `json:"status"`
	Settings string        `json:"settings"`
}

func ToInfoResponse(info *models.Info) *InfoResponse {
	var user *UserResponse

	if info.User != nil {
		userResponse := ToUserResponse(info.User)
		user = &userResponse
	}

	return &InfoResponse{
		ID:       info.ID,
		UserID:   info.UserID,
		User:     user,
		Title:    info.Title,
		Body:     info.Body,
		Status:   info.Status,
		Settings: info.Settings,
	}
}

func ToInfoResponseSlice(info []*models.Info) []*InfoResponse {
	var infoResponse []*InfoResponse

	for _, i := range info {
		info := ToInfoResponse(i)

		infoResponse = append(infoResponse, info)
	}

	return infoResponse
}
