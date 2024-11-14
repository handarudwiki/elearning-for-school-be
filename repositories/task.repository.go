package repositories

import (
	"context"

	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/models"
	"github.com/handarudwiki/models/dto"
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

func (t *taskRepository) Update(ctx context.Context, id int, task *models.Task) (*models.Task, error) {
	if err := t.db.WithContext(ctx).Model(&models.Task{}).Where("id = ?", id).Updates(task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (t *taskRepository) Delete(ctx context.Context, id int) error {
	if err := t.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Task{}).Error; err != nil {
		return err
	}

	return nil
}

func (t *taskRepository) FindAll(ctx context.Context, query dto.QueryDTO) ([]*models.Task, int, error) {
	var tasks []*models.Task
	var count int64

	if err := t.db.Scopes(helpers.SearchTitle(*query.Search), helpers.FilterIsActive(query.Is_active)).Error; err != nil {
		return nil, 0, err
	}

	if err := t.db.Scopes(helpers.Paginate(query.Page, query.Size), helpers.FilterIsActive(query.Is_active), helpers.SearchTitle(*query.Search)).Preload("User").Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	return tasks, int(count), nil
}
