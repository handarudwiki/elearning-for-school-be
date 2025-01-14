package repositories

import (
	"context"

	"github.com/handarudwiki/models"
	"gorm.io/gorm"
)

type clasroomStudentRepository struct {
	db *gorm.DB
}

func NewClassroomStudent(db *gorm.DB) models.ClassroomStudentRepository {
	return &clasroomStudentRepository{
		db: db,
	}
}

func (r *clasroomStudentRepository) FindByClassroomID(ctx context.Context, classroomID int) ([]*models.ClassroomStudent, error) {
	var classroomStudents []*models.ClassroomStudent
	err := r.db.Where("classroom_id = ?", classroomID).Preload("Student").Find(&classroomStudents).Error
	if err != nil {
		return nil, err
	}

	return classroomStudents, nil
}

func (r *clasroomStudentRepository) Create(ctx context.Context, classroomStudent *models.ClassroomStudent) (*models.ClassroomStudent, error) {
	err := r.db.Create(classroomStudent).Error
	if err != nil {
		return nil, err
	}

	return classroomStudent, nil
}
