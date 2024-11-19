package services

import (
	"context"
	"errors"

	"github.com/handarudwiki/models"
	"github.com/handarudwiki/models/commons"
	"github.com/handarudwiki/models/dto"
	"github.com/handarudwiki/models/response"
	"gorm.io/gorm"
)

type ScheduleService interface {
	CreateSchedule(ctx context.Context, schedule *dto.ScheduleDTO) (*response.ScheduleResponse, error)
	GetScheduleByID(ctx context.Context, id int) (*response.ScheduleResponse, error)
}

type scheduleService struct {
	scheduleRepo            models.ScheduleRepository
	classroomSubjectService ClassroomSubjectService
}

func NewSchedule(scheduleRepo models.ScheduleRepository, classroomSubjectService ClassroomSubjectService) ScheduleService {
	return &scheduleService{
		scheduleRepo:            scheduleRepo,
		classroomSubjectService: classroomSubjectService,
	}
}

func (s *scheduleService) CreateSchedule(ctx context.Context, schedule *dto.ScheduleDTO) (*response.ScheduleResponse, error) {
	_, err := s.classroomSubjectService.FindByID(ctx, int(schedule.ClassroomSubjectID))
	if err != nil {
		return nil, err
	}

	if schedule.EndTime <= schedule.StartTime {
		return nil, commons.ErrInvalidInput
	}

	newSchedule := models.Schedule{
		ClassroomSubjectID: schedule.ClassroomSubjectID,
		Day:                schedule.Day,
		StartTime:          schedule.StartTime,
		EndTime:            schedule.EndTime,
	}

	savedSchedule, err := s.scheduleRepo.Create(ctx, &newSchedule)
	if err != nil {
		return nil, err
	}

	return response.ToScheduleResponse(*savedSchedule), nil
}

func (s *scheduleService) GetScheduleByID(ctx context.Context, id int) (*response.ScheduleResponse, error) {
	schedule, err := s.scheduleRepo.FindByID(ctx, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err != nil {
		return nil, errors.New("schedule not found")
	}

	return response.ToScheduleResponse(*schedule), nil
}
