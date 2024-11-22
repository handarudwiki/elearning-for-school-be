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

type StandartService interface {
	Create(ctx context.Context, dto *dto.StandartDTO) (res *response.StandartResponse, err error)
	FindById(ctx context.Context, id int) (res *response.StandartResponse, err error)
}

type standartService struct {
	standartRepo   models.StandartRepository
	subjectService SubjectService
}

func NewStandart(standartRepo models.StandartRepository, subjectService SubjectService) StandartService {
	return &standartService{
		standartRepo:   standartRepo,
		subjectService: subjectService,
	}
}

func (s *standartService) Create(ctx context.Context, dto *dto.StandartDTO) (res *response.StandartResponse, err error) {
	_, err = s.subjectService.FindByID(ctx, int(dto.SubjectID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, commons.ErrNotFound
	}

	standart := &models.Standart{
		SubjectID:  dto.SubjectID,
		Type:       dto.Type,
		Body:       dto.Body,
		Code:       dto.Code,
		StandartID: dto.StandartID,
		TeacherID:  dto.TeacherID,
	}

	standart, err = s.standartRepo.Create(ctx, standart)
	if err != nil {
		return nil, err
	}

	res = response.ToStandartResponse(standart)
	return res, nil
}

func (s *standartService) FindById(ctx context.Context, id int) (res *response.StandartResponse, err error) {
	standart, err := s.standartRepo.FindByID(ctx, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, commons.ErrNotFound
	}

	res = &response.StandartResponse{
		ID:        standart.ID,
		SubjectID: standart.SubjectID,
		Type:      standart.Type,
		Body:      standart.Body,
		Code:      standart.Code,
	}

	return res, nil
}
