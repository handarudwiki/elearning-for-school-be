package dto

type CreateClassroomDTO struct {
	TeacherID uint   `json:"teacher_id" validate:"required,gt=0"`
	Name      string `json:"name" validate:"required"`
	Grade     string `json:"grade" validate:"required"`
	Group     string `json:"group" validate:"required"`
	Settings  string `json:"settings" validate:"required"`
}

type UpdateClassroomDTO struct {
	TeacherID uint   `json:"teacher_id" validate:"required,gt=0"`
	Name      string `json:"name" validate:"required"`
	Grade     string `json:"grade" validate:"required"`
	Group     string `json:"group" validate:"required"`
	Settings  string `json:"settings" validate:"required"`
}
