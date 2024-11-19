package response

import (
	"github.com/handarudwiki/models"
)

type ScheduleResponse struct {
	ID                 uint                      `json:"id"`
	ClassroomSubjectID uint                      `json:"classroom_id"`
	ClassroomSubject   *ClassroomSubjectResponse `json:"classroom_subject,omitempty"`
	Day                int                       `json:"day"`
	StartTime          string                    `json:"start_at"`
	EndTime            string                    `json:"end_at"`
}

func ToScheduleResponse(schedule models.Schedule) *ScheduleResponse {
	var classroomSubject *ClassroomSubjectResponse

	if schedule.ClassroomSubject != nil {
		classroomSubjectResp := ToClassroomSubjectResponse(schedule.ClassroomSubject)
		classroomSubject = &classroomSubjectResp
	}

	return &ScheduleResponse{
		ID:                 schedule.ID,
		ClassroomSubjectID: schedule.ClassroomSubjectID,
		ClassroomSubject:   classroomSubject,
		Day:                schedule.Day,
		StartTime:          schedule.StartTime,
		EndTime:            schedule.EndTime,
	}
}

func ToScheduleResponsesSlice(schedules []*models.Schedule) []*ScheduleResponse {
	var scheduleResponses []*ScheduleResponse
	for _, schedule := range schedules {
		scheduleResponses = append(scheduleResponses, ToScheduleResponse(*schedule))
	}
	return scheduleResponses
}
