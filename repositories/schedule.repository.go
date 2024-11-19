package repositories

import (
	"context"

	"github.com/handarudwiki/models"
	"gorm.io/gorm"
)

type scheduleRepository struct {
	db *gorm.DB
}

func NewSchedule(db *gorm.DB) models.ScheduleRepository {
	return &scheduleRepository{
		db: db,
	}
}

func (r *scheduleRepository) FindByID(ctx context.Context, id int) (res *models.Schedule, err error) {
	var schedule models.Schedule
	err = r.db.Where("id = ?", id).First(&schedule).Error

	if err != nil {
		return nil, err
	}

	return &schedule, nil
}

func (r *scheduleRepository) Create(ctx context.Context, schedule *models.Schedule) (*models.Schedule, error) {

	err := r.db.Create(schedule).Error
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

func (r *scheduleRepository) FindByClassroomSubjectID(ctx context.Context, classroomSubjectID int) (res []*models.Schedule, err error) {
	var schedule []*models.Schedule
	err = r.db.Where("classroom_subject_id = ?", classroomSubjectID).Order("start_time ").Find(&schedule).Error

	if err != nil {
		return nil, err
	}

	return schedule, nil
}
