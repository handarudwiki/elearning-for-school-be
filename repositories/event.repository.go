package repositories

import (
	"context"

	"github.com/handarudwiki/models"
	"gorm.io/gorm"
)

type eventRepository struct {
	db *gorm.DB
}

func NewEvent(db *gorm.DB) models.EventRepository {
	return &eventRepository{
		db: db,
	}
}

func (r *eventRepository) FindByID(ctx context.Context, id int) (res *models.Event, err error) {
	var event models.Event
	err = r.db.Where("id = ?", id).Preload("User").First(&event).Error

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *eventRepository) Create(ctx context.Context, event *models.Event) (*models.Event, error) {

	err := r.db.Create(event).Error
	if err != nil {
		return nil, err
	}

	return event, nil
}
