package repositories

import (
	"context"

	"github.com/handarudwiki/models"
	"gorm.io/gorm"
)

type classroomRepository struct {
	db *gorm.DB
}

func NewClassroom(db *gorm.DB) models.ClassroomRepository {
	return &classroomRepository{
		db: db,
	}
}

func (r *classroomRepository) FindByID(ctx context.Context, id int) (res *models.Classroom, err error) {
	var classroom models.Classroom
	err = r.db.Where("id = ?", id).First(&classroom).Error
	if err != nil {
		return nil, err
	}

	return &classroom, nil
}

func (r *classroomRepository) Create(ctx context.Context, classroom *models.Classroom) (*models.Classroom, error) {
	err := r.db.Create(classroom).Error
	if err != nil {
		return nil, err
	}

	return classroom, nil
}
