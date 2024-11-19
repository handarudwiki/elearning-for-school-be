package dto

type ScheduleDTO struct {
	ID                 uint   `json:"id"`
	ClassroomSubjectID uint   `json:"classroom_subject_id" validate:"required"`
	Day                int    `json:"day" validate:"required,gt=0,lte=7"`
	StartTime          string `json:"start_time" validate:"required,time"`
	EndTime            string `json:"end_time" validate:"required,time"`
}
