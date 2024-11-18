package response

import "github.com/handarudwiki/models"

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Details  string `json:"details"`
	IsActive bool   `json:"is_active"`
	Role     int    `json:"role"`
}

func ToUserResponse(user *models.User) UserResponse {
	return UserResponse{
		ID:       int(user.ID),
		Name:     user.Name,
		Email:    user.Email,
		Details:  user.Details,
		IsActive: user.IsActive,
		Role:     user.Role,
	}

}

func ToUserResponseSlice(user []*models.User) []UserResponse {
	var userResponses []UserResponse
	for _, u := range user {
		userResponses = append(userResponses, ToUserResponse(u))
	}
	return userResponses
}
