package response

import "github.com/handarudwiki/models"

type LectureResponse struct {
	ID        uint            `json:"id"`
	SubjectID uint            `json:"subject_id"`
	Subject   *models.Subject `json:"subject,omitempty"`
	UserID    uint            `json:"user_id"`
	User      *models.User    `json:"user,omitempty"`
	Title     string          `json:"title"`
	Body      string          `json:"body"`
	IsActive  bool            `json:"is_active"`
	Addition  string          `json:"additions"`
}

func ToLectureResponse(lecture *models.Lecture) LectureResponse {
	return LectureResponse{
		ID:        lecture.ID,
		SubjectID: lecture.SubjectID,
		Title:     lecture.Title,
		Body:      lecture.Body,
		IsActive:  lecture.IsActive,
		Addition:  lecture.Addition,
		Subject:   &lecture.Subject,
		User:      &lecture.User,
	}
}

func ToLectureResponseSlice(lectures []*models.Lecture) []LectureResponse {
	var lectureResponses []LectureResponse
	for _, lecture := range lectures {
		lectureResponses = append(lectureResponses, ToLectureResponse(lecture))
	}
	return lectureResponses
}
