package repositories

import (
	"context"

	"github.com/handarudwiki/models"
	"gorm.io/gorm"
)

type classroomTaskRepository struct {
	db *gorm.DB
}

func NewClassroomTask(db *gorm.DB) models.ClassroomTaskRepository {
	return &classroomTaskRepository{
		db: db,
	}
}

func (r *classroomTaskRepository) Create(ctx context.Context, classroomTask *models.ClassroomTask) (*models.ClassroomTask, error) {
	err := r.db.Create(classroomTask).Error
	if err != nil {
		return nil, err
	}

	return classroomTask, nil
}

func (r *classroomTaskRepository) FindByID(ctx context.Context, id int) (res *models.ClassroomTask, err error) {
	var classroomTask models.ClassroomTask
	err = r.db.Where("id = ?", id).Preload("Task").Preload("Teacher").Preload("Classroom").First(&classroomTask).Error
	if err != nil {
		return nil, err
	}

	return &classroomTask, nil
}

func (r *classroomTaskRepository) FindByClassroomID(ctx context.Context, classroomID int) (res []*models.ClassroomTask, err error) {
	var classroomTask []*models.ClassroomTask
	err = r.db.Where("classroom_id = ?", classroomID).Preload("Task").Preload("Teacher").Preload("Classroom").Find(&classroomTask).Error
	if err != nil {
		return nil, err
	}

	return classroomTask, nil
}
