package services

import (
	"context"
	"errors"

	"github.com/handarudwiki/models"
	"github.com/handarudwiki/models/commons"
	"github.com/handarudwiki/models/dto"
	"github.com/handarudwiki/models/response"
	"gorm.io/gorm"
)

type EventService interface {
	FindById(ctx context.Context, id int) (res *response.EventResponse, err error)
	Create(ctx context.Context, dto *dto.EventDto) (res *response.EventResponse, err error)
}

type eventService struct {
	eventRepo models.EventRepository
}

func NewEvent(eventRepo models.EventRepository) EventService {
	return &eventService{
		eventRepo: eventRepo,
	}
}

func (s *eventService) Create(ctx context.Context, dto *dto.EventDto) (res *response.EventResponse, err error) {

	if err != nil {
		return res, err
	}

	event := &models.Event{
		Title:    dto.Title,
		Body:     dto.Body,
		UserID:   dto.UserID,
		Location: dto.Location,
		Date:     dto.Date,
		Time:     dto.Time,
	}

	event, err = s.eventRepo.Create(ctx, event)

	if err != nil {
		return res, err
	}

	res = response.ToEventResponse(event)

	return res, nil
}

func (s *eventService) FindById(ctx context.Context, id int) (res *response.EventResponse, err error) {
	event, err := s.eventRepo.FindByID(ctx, id)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return res, err

	}

	if err != nil {
		return res, commons.ErrNotFound
	}

	res = response.ToEventResponse(event)

	return res, nil
}
