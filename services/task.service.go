package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/handarudwiki/models"
	"github.com/handarudwiki/models/commons"
	"github.com/handarudwiki/models/dto"
	"github.com/handarudwiki/models/response"
	"gorm.io/gorm"
)

type TaskService interface {
	Create(ctx context.Context, dto dto.CreateTaskDTO) (response.TaskResponse, error)
	FindByID(ctx context.Context, id int) (response.TaskResponse, error)
}

type taskService struct {
	taskRepository models.TaskRepository
}

func NewTask(taskRepository models.TaskRepository, userService UserService) TaskService {
	return &taskService{
		taskRepository: taskRepository,
	}
}

func (s *taskService) Create(ctx context.Context, dto dto.CreateTaskDTO) (response.TaskResponse, error) {

	task := models.Task{
		Title:    dto.Title,
		UserID:   dto.UserID,
		Type:     dto.Type,
		Body:     dto.Body,
		IsActive: dto.IsActive,
		Settings: dto.Settings,
		Deadline: dto.Deadline,
	}

	fmt.Println(task)

	newTask, err := s.taskRepository.Create(ctx, &task)
	if err != nil {
		return response.TaskResponse{}, err
	}

	res := response.ToTaskResponse(*newTask)

	return res, nil
}

func (s *taskService) FindByID(ctx context.Context, id int) (response.TaskResponse, error) {
	task, err := s.taskRepository.FindByID(ctx, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return response.TaskResponse{}, err
	}

	if err != nil {
		return response.TaskResponse{}, commons.ErrNotFound
	}

	res := response.ToTaskResponse(*task)

	return res, nil
}
