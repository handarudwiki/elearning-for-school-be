package repositories

import (
	"context"
	"math"

	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/models"
	"github.com/handarudwiki/models/dto"
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

func (r *eventRepository) Update(ctx context.Context, event *models.Event, id int) (*models.Event, error) {
	err := r.db.Model(&models.Event{}).Where("id = ?", id).Updates(event).Error
	if err != nil {
		return nil, err
	}

	return event, nil
}

func (r *eventRepository) Delete(ctx context.Context, id int) error {
	err := r.db.Where("id = ?", id).Delete(&models.Event{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *eventRepository) FindAll(ctx context.Context, dto dto.QueryDTO) ([]*models.Event, int, error) {
	var events []*models.Event
	err := r.db.Preload("User").
		Scopes(helpers.Paginate(dto.Page, dto.Size), helpers.SearchTitle(*dto.Search)).
		Find(&events).Error

	if err != nil {
		return nil, 0, err
	}

	var total int64

	err = r.db.Model(events).
		Scopes(helpers.SearchTitle(*dto.Search)).
		Count(&total).Error

	totalPage := math.Ceil(float64(total) / float64(dto.Size))

	if err != nil {
		return nil, 0, err
	}

	return events, int(totalPage), nil

}
