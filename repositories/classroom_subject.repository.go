package repositories

import (
	"context"

	"github.com/handarudwiki/models"
	"gorm.io/gorm"
)

type classroomSubjectRepository struct {
	db *gorm.DB
}

func NewClassroomSubject(db *gorm.DB) models.ClassroomSubjectRepository {
	return &classroomSubjectRepository{
		db: db,
	}
}

func (r *classroomSubjectRepository) Create(ctx context.Context, classroomSubject *models.ClassroomSubject) (*models.ClassroomSubject, error) {
	err := r.db.Create(classroomSubject).Error
	if err != nil {
		return nil, err
	}

	return classroomSubject, nil
}

func (r *classroomSubjectRepository) FindByTeacherID(ctx context.Context, teacherID int) ([]*models.ClassroomSubject, error) {
	var classroomSubjects []*models.ClassroomSubject
	err := r.db.Where("teacher_id = ?", teacherID).Preload("Classroom").Preload("Subject").Preload("Teacher").Find(&classroomSubjects).Error
	if err != nil {
		return nil, err
	}

	return classroomSubjects, nil
}