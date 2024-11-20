package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/handarudwiki/models"
	"github.com/handarudwiki/models/commons"
	"github.com/handarudwiki/models/dto"
	"github.com/handarudwiki/models/response"
	"gorm.io/gorm"
)

type ScheduleService interface {
	CreateSchedule(ctx context.Context, schedule *dto.ScheduleDTO) (*response.ScheduleResponse, error)
	GetScheduleByID(ctx context.Context, id int) (*response.ScheduleResponse, error)
	GetByClassroomSubjectID(ctx context.Context, classroomSubjectID int) ([]*response.ScheduleResponse, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, schedule *dto.ScheduleDTO, id int) (*response.ScheduleResponse, error)
	GetScheduleByday(ctx context.Context, teacherID int) ([]map[string]interface{}, error)
	GetScheduleClassroomDay(ctx context.Context, classroomID int, teacherId *int) ([]map[string]interface{}, error)
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

func (s *scheduleService) GetByClassroomSubjectID(ctx context.Context, classroomSubjectID int) ([]*response.ScheduleResponse, error) {
	schedules, err := s.scheduleRepo.FindByClassroomSubjectID(ctx, classroomSubjectID)
	if err != nil {
		return nil, err
	}

	res := response.ToScheduleResponsesSlice(schedules)

	return res, nil
}

func (s *scheduleService) Delete(ctx context.Context, id int) error {
	_, err := s.scheduleRepo.FindByID(ctx, id)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err != nil {
		return commons.ErrNotFound
	}
	err = s.scheduleRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *scheduleService) Update(ctx context.Context, schedule *dto.ScheduleDTO, id int) (*response.ScheduleResponse, error) {
	_, err := s.classroomSubjectService.FindByID(ctx, int(schedule.ClassroomSubjectID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err != nil {
		return nil, commons.ErrNotFound
	}

	if schedule.EndTime <= schedule.StartTime {
		return nil, commons.ErrInvalidInput
	}

	scheduleToUpdate := models.Schedule{
		ClassroomSubjectID: schedule.ClassroomSubjectID,
		Day:                schedule.Day,
		StartTime:          schedule.StartTime,
		EndTime:            schedule.EndTime,
	}

	updatedSchedule, err := s.scheduleRepo.Update(ctx, id, &scheduleToUpdate)
	if err != nil {
		return nil, err
	}

	return response.ToScheduleResponse(*updatedSchedule), nil
}

func (s *scheduleService) GetScheduleByday(ctx context.Context, teacherID int) ([]map[string]interface{}, error) {

	dayOfWeek := int(time.Now().Weekday())

	// fmt.Println("hari ini ", dayOfWeek)

	schedules, err := s.scheduleRepo.GetScheduleByday(ctx, dayOfWeek, teacherID)
	if err != nil {
		return nil, err
	}

	fmt.Println(schedules)

	var res []map[string]interface{}

	for _, schedule := range schedules {
		item := map[string]interface{}{
			"id":             schedule.ID,
			"classroom_name": schedule.ClassroomSubject.Classroom.Name,
			"subject_name":   schedule.ClassroomSubject.Subject.Name,
			"start_time":     schedule.StartTime,
			"end_time":       schedule.EndTime,
		}
		res = append(res, item)
	}

	return res, nil
}

func (s *scheduleService) GetScheduleClassroomDay(ctx context.Context, classroomID int, teacherId *int) (res []map[string]interface{}, err error) {
	var schedules []*models.Schedule

	dayOfWeek := int(time.Now().Weekday())

	fmt.Println("hari ini ", dayOfWeek)

	if *teacherId == 1 {
		schedules, err = s.scheduleRepo.GetdataSchedulesClassroomDay(ctx, dayOfWeek, classroomID, teacherId)
	} else {
		schedules, err = s.scheduleRepo.GetdataSchedulesClassroomDay(ctx, dayOfWeek, classroomID, nil)
	}
	if err != nil {
		return nil, err
	}

	for _, schedule := range schedules {
		item := map[string]interface{}{
			"id":             schedule.ID,
			"classroom_name": schedule.ClassroomSubject.Classroom.Name,
			"subject_name":   schedule.ClassroomSubject.Subject.Name,
			"start_time":     schedule.StartTime,
			"end_time":       schedule.EndTime,
		}

		res = append(res, item)
	}

	return res, nil
}
