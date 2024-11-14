package repositories

import (
	"context"

	"github.com/handarudwiki/models"
	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

func NewTask(db *gorm.DB) models.TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (t *taskRepository) Create(ctx context.Context, task *models.Task) (*models.Task, error) {
	if err := t.db.WithContext(ctx).Create(task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (t *taskRepository) FindByID(ctx context.Context, id int) (*models.Task, error) {
	var task models.Task
	if err := t.db.WithContext(ctx).Where("id = ?", id).Preload("User").First(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}
