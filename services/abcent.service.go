package services

import (
	"context"

	"github.com/handarudwiki/models"
	"github.com/handarudwiki/models/commons"
	"github.com/handarudwiki/models/dto"
	"github.com/handarudwiki/models/response"
)

type AbcentService interface {
	Create(ctx context.Context, dto dto.CreateAbcentDTO) (res *response.AbcentResponse, err error)
	FindByScheduleIDToday(ctx context.Context, scheduleID int, date string, query dto.QueryDTO) (res []*response.AbcentResponse, paginate commons.Paginate, err error)
}

type abcentService struct {
	abcentRepo models.AbcentRepository
}

func NewAbcent(abcentRepo models.AbcentRepository) AbcentService {
	return &abcentService{
		abcentRepo: abcentRepo,
	}
}

func (s *abcentService) Create(ctx context.Context, dto dto.CreateAbcentDTO) (res *response.AbcentResponse, err error) {

	abcent := &models.Abcent{
		UserID:      uint(dto.UserID),
		ScheduleID:  uint(dto.ScheduleID),
		IsAbcent:    dto.IsAbcent,
		Reason:      dto.Reason,
		Description: dto.Description,
		Details:     dto.Details,
	}

	abcent, err = s.abcentRepo.Create(ctx, abcent)

	if err != nil {
		return res, err
	}

	res = response.ToAbcentResponse(abcent)

	return res, nil
}

func (s *abcentService) FindByScheduleIDToday(ctx context.Context, scheduleID int, date string, query dto.QueryDTO) (res []*response.AbcentResponse, paginate commons.Paginate, err error) {
	abcents, count, err := s.abcentRepo.FindByScheduleIDToday(ctx, scheduleID, date, query)
	if err != nil {
		return
	}

	paginate = commons.ToPaginate(query.Page, query.Size, int(count))

	res = response.ToAbcentResponses(abcents)

	return
}
