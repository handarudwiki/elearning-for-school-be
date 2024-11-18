package response

import (
	"time"

	"github.com/handarudwiki/models"
)

type TaskResponse struct {
	ID       uint         `json:"id"`
	Title    string       `json:"title"`
	UserID   uint         `json:"user_id"`
	User     *models.User `json:"user"`
	Type     int          `json:"type"`
	Body     string       `json:"body"`
	IsActive bool         `json:"is_active"`
	Settings string       `json:"settings"`
	Deadline time.Time    `json:"deadline"`
}

func ToTaskResponse(task *models.Task) TaskResponse {
	return TaskResponse{
		ID:       task.ID,
		Title:    task.Title,
		UserID:   task.UserID,
		User:     task.User,
		Type:     task.Type,
		Body:     task.Body,
		IsActive: task.IsActive,
		Settings: task.Settings,
		Deadline: task.Deadline,
	}
}

func ToTaskResponseSlice(task []*models.Task) []TaskResponse {
	var taskResponseSlice []TaskResponse
	for _, task := range task {
		taskResponseSlice = append(taskResponseSlice, ToTaskResponse(task))
	}
	return taskResponseSlice
}
