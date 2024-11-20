package services

import (
	"context"
	"errors"

	"github.com/handarudwiki/models"
	"github.com/handarudwiki/models/dto"
	"github.com/handarudwiki/models/response"
	"gorm.io/gorm"
)

type InfoService interface {
	FindById(ctx context.Context, id int) (res *response.InfoResponse, err error)
	Create(ctx context.Context, info *dto.InfoDto) (res *response.InfoResponse, err error)
}

type infoService struct {
	infoRepo models.InfoRepository
}

func NewInfo(infoRepo models.InfoRepository) InfoService {
	return &infoService{
		infoRepo: infoRepo,
	}
}

func (s *infoService) Create(ctx context.Context, dto *dto.InfoDto) (res *response.InfoResponse, err error) {

	info := &models.Info{
		Title:    dto.Title,
		Body:     dto.Body,
		UserID:   int(dto.UserID),
		Status:   dto.Status,
		Settings: dto.Settings,
	}

	info, err = s.infoRepo.Create(ctx, info)

	if err != nil {
		return res, err
	}

	res = response.ToInfoResponse(info)

	return res, nil
}

func (s *infoService) FindById(ctx context.Context, id int) (res *response.InfoResponse, err error) {
	info, err := s.infoRepo.FindByID(ctx, id)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return res, err
	}

	res = response.ToInfoResponse(info)

	return res, nil
}
