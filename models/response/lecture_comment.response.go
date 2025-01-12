package response

import "github.com/handarudwiki/models"

type LectureCommentResponse struct {
	ID        uint          `json:"id"`
	LectureID uint          `json:"lecture_id"`
	Content   string        `json:"content"`
	UserID    uint          `json:"user_id"`
	Users     *UserResponse `json:"user"`
}

func ToLectureCommentResponse(lectureComment *models.LectureComment) LectureCommentResponse {
	var userResponse UserResponse

	if lectureComment.User != nil {
		userResponse = ToUserResponse(lectureComment.User)
	}
	return LectureCommentResponse{
		ID:        lectureComment.ID,
		LectureID: lectureComment.LectureID,
		Content:   lectureComment.Content,
		UserID:    lectureComment.UserID,
		Users:     &userResponse,
	}
}

func ToLectureCommentResponseSlice(lectureComments []*models.LectureComment) []LectureCommentResponse {
	var lectureCommentResponses []LectureCommentResponse
	for _, lectureComment := range lectureComments {
		lectureCommentResponses = append(lectureCommentResponses, ToLectureCommentResponse(lectureComment))
	}
	return lectureCommentResponses
}
