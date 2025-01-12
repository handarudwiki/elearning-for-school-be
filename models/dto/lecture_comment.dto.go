package dto

type CreateLectureCommentDTO struct {
	LectureID uint   `json:"lecture_id" validate:"required"`
	Content   string `json:"content" validate:"required"`
	UserID    uint   `json:"-"`
}
