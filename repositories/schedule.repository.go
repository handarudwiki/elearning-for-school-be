package repositories

import (
	"context"
	"fmt"

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

func (r *scheduleRepository) Delete(ctx context.Context, id int) error {
	err := r.db.Where("id = ?", id).Delete(&models.Schedule{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *scheduleRepository) Update(ctx context.Context, id int, schedule *models.Schedule) (*models.Schedule, error) {
	err := r.db.Where("id = ?", id).Updates(schedule).Error
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

func (r *scheduleRepository) GetScheduleByday(ctx context.Context, day, teacherID int) (res []*models.Schedule, err error) {
	var schedule []*models.Schedule
	err = r.db.Where("day = ?", day).
		Preload("ClassroomSubject.Classroom").
		Preload("ClassroomSubject.Subject").
		Where("classroom_subject_id IN (SELECT id from classroom_subjects WHERE teacher_id =? )", teacherID).Order("start_time").
		Find(&schedule).Error

	if err != nil {
		return nil, err
	}

	return schedule, nil
}

func (r *scheduleRepository) GetdataSchedulesClassroomDay(ctx context.Context, day, classroomID int, teacherId *int) ([]*models.Schedule, error) {
	var schedules []*models.Schedule

	fmt.Println("classroom ", classroomID)

	subQuery := r.db.Model(&models.ClassroomSubject{}).Select("id").
		Where("classroom_id=?", classroomID)

	if teacherId != nil {
		subQuery = subQuery.Where("teacher_id = ?", *teacherId)
	}

	err := r.db.Preload("ClassroomSubject.Classroom").
		Preload("ClassroomSubject.Subject").
		Where("day = ?", day).
		Where("classroom_subject_id IN (?) ", subQuery).Order("start_time").Find(&schedules).Error

	if err != nil {
		return nil, err
	}

	return schedules, nil
}
