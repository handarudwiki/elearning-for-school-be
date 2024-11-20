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
	Update(ctx context.Context, dto *dto.EventDto, id int) (res *response.EventResponse, err error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context, query dto.QueryDTO) (res []*response.EventResponse, page commons.Paginate, err error)
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

func (s *eventService) Update(ctx context.Context, dto *dto.EventDto, id int) (res *response.EventResponse, err error) {

	_, err = s.eventRepo.FindByID(ctx, id)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return res, err
	}

	if err != nil {
		return res, commons.ErrNotFound
	}

	event := &models.Event{
		Title:    dto.Title,
		Body:     dto.Body,
		UserID:   dto.UserID,
		Location: dto.Location,
		Date:     dto.Date,
		Time:     dto.Time,
	}

	event, err = s.eventRepo.Update(ctx, event, id)

	if err != nil {
		return res, err
	}

	res = response.ToEventResponse(event)

	return res, nil
}

func (s *eventService) Delete(ctx context.Context, id int) error {
	_, err := s.eventRepo.FindByID(ctx, id)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err != nil {
		return commons.ErrNotFound
	}

	err = s.eventRepo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *eventService) FindAll(ctx context.Context, query dto.QueryDTO) (res []*response.EventResponse, page commons.Paginate, err error) {
	events, total, err := s.eventRepo.FindAll(ctx, query)

	if err != nil {
		return res, page, err
	}

	res = response.ToEventResponseSlice(events)

	page = commons.Paginate{
		TotalPage: total,
		Page:      query.Page,
		Size:      query.Size,
	}

	return res, page, nil
}
