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
	Update(ctx context.Context, dto *dto.StandartDTO, id int) (res *response.StandartResponse, err error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context, query *dto.QueryDTO) (res []*response.StandartResponse, page commons.Paginate, err error)
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

	res = response.ToStandartResponse(standart)

	return res, nil
}

func (s *standartService) Update(ctx context.Context, dto *dto.StandartDTO, id int) (res *response.StandartResponse, err error) {
	standart, err := s.standartRepo.FindByID(ctx, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, commons.ErrNotFound
	}

	standart.SubjectID = dto.SubjectID
	standart.Type = dto.Type
	standart.Body = dto.Body
	standart.Code = dto.Code
	standart.StandartID = dto.StandartID
	standart.TeacherID = dto.TeacherID

	standart, err = s.standartRepo.Update(ctx, standart, id)
	if err != nil {
		return nil, err
	}
	standart.ID = uint(id)

	res = response.ToStandartResponse(standart)

	return res, nil
}

func (s *standartService) Delete(ctx context.Context, id int) error {
	_, err := s.standartRepo.FindByID(ctx, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return commons.ErrNotFound
	}

	err = s.standartRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *standartService) FindAll(ctx context.Context, query *dto.QueryDTO) (res []*response.StandartResponse, page commons.Paginate, err error) {
	standarts, total, err := s.standartRepo.FindAll(ctx, query)
	if err != nil {
		return nil, page, err
	}

	res = response.ToStandartResponseSlice(standarts)

	page = commons.ToPaginate(
		query.Page,
		query.Size,
		total,
	)

	return res, page, nil
}
