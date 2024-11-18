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

type ClassroomTaskService interface {
	AsignTask(ctx context.Context, dto dto.AsignTaskDTO) (res []response.ClassroomTaskResponse, err error)
	FindByID(ctx context.Context, id int) (res response.ClassroomTaskResponse, err error)
}

type classroomTaskService struct {
	classroomTaskRepository models.ClassroomTaskRepository
	taskService             TaskService
	clasroomService         ClassroomService
}

func NewClassroomTask(classroomTaskRepository models.ClassroomTaskRepository, taskService TaskService, clasroomService ClassroomService) ClassroomTaskService {
	return &classroomTaskService{
		classroomTaskRepository: classroomTaskRepository,
		taskService:             taskService,
		clasroomService:         clasroomService,
	}
}

func (s *classroomTaskService) AsignTask(ctx context.Context, dto dto.AsignTaskDTO) (res []response.ClassroomTaskResponse, err error) {

	for _, classroomID := range dto.ClassroomIDS {
		_, err := s.clasroomService.FindByID(ctx, int(classroomID))
		if err != nil {
			return res, err
		}
	}

	_, err = s.taskService.FindByID(ctx, int(dto.TaskID))
	if err != nil {
		return res, err
	}

	for _, classroomID := range dto.ClassroomIDS {
		newClassroomTask := models.ClassroomTask{
			ClassroomID: classroomID,
			TaskID:      dto.TaskID,
			Body:        dto.Body,
			TeacherID:   dto.TeacherID,
		}
		taskClassroom, err := s.classroomTaskRepository.Create(ctx, &newClassroomTask)

		if err != nil {
			return res, err
		}

		res = append(res, response.ToClassroomTaskResponse(taskClassroom))
	}
	return res, nil
}

func (s *classroomTaskService) FindByID(ctx context.Context, id int) (res response.ClassroomTaskResponse, err error) {
	classroomTask, err := s.classroomTaskRepository.FindByID(ctx, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return res, err
	}

	if err != nil {
		return res, commons.ErrNotFound
	}

	res = response.ToClassroomTaskResponse(classroomTask)

	return res, nil

}
