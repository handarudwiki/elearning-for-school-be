package repositories

import (
	"context"

	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/models"
	"github.com/handarudwiki/models/dto"
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
	err = r.db.Where("id = ?", id).Preload("Teacher").First(&classroom).Error
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

func (r *classroomRepository) Update(ctx context.Context, classroom *models.Classroom, id int) (*models.Classroom, error) {
	err := r.db.Model(&models.Classroom{}).Where("id = ?", id).Updates(classroom).Error
	if err != nil {
		return nil, err
	}

	return classroom, nil
}

func (r *classroomRepository) Delete(ctx context.Context, id int) error {
	err := r.db.Where("id = ?", id).Delete(&models.Classroom{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *classroomRepository) FindAll(ctx context.Context, dto dto.QueryDTO) (classrooms []*models.Classroom, totalPage int64, err error) {

	err = r.db.Model(&models.Classroom{}).Scopes(helpers.Search(*dto.Search)).Count(
		&totalPage,
	).Error

	if err != nil {
		return
	}

	err = r.db.Scopes(helpers.Paginate(dto.Page, dto.Size), helpers.Search(*dto.Search)).Preload("Teacher").Find(&classrooms).Error

	if err != nil {
		return nil, totalPage, err
	}

	return classrooms, totalPage, nil
}
