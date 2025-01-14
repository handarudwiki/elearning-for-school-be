package repositories

import (
	"context"
	"time"

	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/models"
	"github.com/handarudwiki/models/dto"
	"gorm.io/gorm"
)

type abcentRepository struct {
	db *gorm.DB
}

func NewAbcent(db *gorm.DB) models.AbcentRepository {
	return &abcentRepository{
		db: db,
	}
}

func (r *abcentRepository) Create(ctx context.Context, abcent *models.Abcent) (*models.Abcent, error) {
	err := r.db.Create(abcent).Error
	if err != nil {
		return nil, err
	}

	return abcent, nil
}

func (r *abcentRepository) FindByScheduleIDToday(ctx context.Context, scheduleID int, date string, query dto.QueryDTO) (abcents []*models.Abcent, count int64, err error) {
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	err = r.db.Debug().Preload("User").
		Scopes(helpers.Paginate(query.Page, query.Size)).
		Where("schedule_id", scheduleID).
		Where("DATE(abcents.created_at) = ?", date).
		Joins("Join users on users.id = abcents.user_id AND users.role = ?", "2").
		Find(&abcents).Error

	if err != nil {
		return
	}

	err = r.db.Model(&models.Abcent{}).
		Where("schedule_id", scheduleID).
		Where("DATE(abcents.created_at) = ?", date).
		Joins("Join users on users.id = abcents.user_id AND users.role = ?", "2").
		Count(&count).Error

	return
}

func (r *abcentRepository) Update(ctx context.Context, abcent *models.Abcent, id int) (*models.Abcent, error) {
	err := r.db.Model(abcent).Where("id = ?", id).Updates(abcent).Update("is_abcent", abcent.IsAbcent).Error
	if err != nil {
		return nil, err
	}

	return abcent, nil
}
