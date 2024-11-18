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
	Update(ctx context.Context, id int, dto dto.UpdateTaskDTO) (response.TaskResponse, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context, query dto.QueryDTO) ([]response.TaskResponse, commons.Paginate, error)
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

	res := response.ToTaskResponse(newTask)

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

	res := response.ToTaskResponse(task)

	return res, nil
}

func (s *taskService) Update(ctx context.Context, id int, dto dto.UpdateTaskDTO) (response.TaskResponse, error) {
	task, err := s.taskRepository.FindByID(ctx, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return response.TaskResponse{}, err
	}

	if err != nil {
		return response.TaskResponse{}, commons.ErrNotFound
	}

	task.Title = dto.Title
	task.User.ID = dto.UserID
	task.Type = dto.Type
	task.Body = dto.Body
	task.IsActive = dto.IsActive
	task.Settings = dto.Settings
	task.Deadline = dto.Deadline

	newTask, err := s.taskRepository.Update(ctx, id, task)
	if err != nil {
		return response.TaskResponse{}, err
	}

	fmt.Println(newTask)
	res := response.ToTaskResponse(newTask)

	return res, nil
}

func (s *taskService) Delete(ctx context.Context, id int) error {

	_, err := s.taskRepository.FindByID(ctx, id)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err != nil {
		return commons.ErrNotFound
	}

	err = s.taskRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *taskService) GetAll(ctx context.Context, query dto.QueryDTO) ([]response.TaskResponse, commons.Paginate, error) {
	tasks, totalPage, err := s.taskRepository.FindAll(ctx, query)
	if err != nil {
		return nil, commons.Paginate{}, err
	}

	page := commons.ToPaginate(query.Page, query.Size, totalPage)

	res := response.ToTaskResponseSlice(tasks)

	return res, page, nil
}
