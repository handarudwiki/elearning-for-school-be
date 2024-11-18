package dto

type AsignTaskDTO struct {
	TaskID       uint   `json:"task_id" validate:"required,gt=0"`
	ClassroomIDS []uint `json:"classroom_id" validate:"required,gt=0"`
	Body         string `json:"body" validate:"required"`
	TeacherID    uint   `json:"-"`
}
