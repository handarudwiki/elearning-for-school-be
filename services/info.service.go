package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/handarudwiki/models"
	"github.com/handarudwiki/models/commons"
	"github.com/handarudwiki/models/dto"
	"github.com/handarudwiki/models/response"
	"gorm.io/gorm"
)

type InfoService interface {
	FindById(ctx context.Context, id int) (res *response.InfoResponse, err error)
	Create(ctx context.Context, info *dto.InfoDto) (res *response.InfoResponse, err error)
	Update(ctx context.Context, info *dto.InfoDto, id int) (res *response.InfoResponse, err error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context, query dto.QueryDTO) (res []*response.InfoResponse, page commons.Paginate, err error)
	FindPublicInfo(ctx context.Context, query dto.QueryDTO) (res []*response.InfoResponse, page commons.Paginate, err error)
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

func (s *infoService) Update(ctx context.Context, dto *dto.InfoDto, id int) (res *response.InfoResponse, err error) {
	info := &models.Info{
		Title:    dto.Title,
		Body:     dto.Body,
		UserID:   int(dto.UserID),
		Status:   dto.Status,
		Settings: dto.Settings,
	}

	info, err = s.infoRepo.Update(ctx, id, info)

	if err != nil {
		return res, err
	}

	fmt.Println(info.ID)

	res = response.ToInfoResponse(info)

	return res, nil
}

func (s *infoService) Delete(ctx context.Context, id int) error {
	// _, err := s.infoRepo.FindByID(ctx, id)

	// if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return err
	// }

	// if err != nil {
	// 	return commons.ErrNotFound
	// }

	err := s.infoRepo.Delete(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (s *infoService) FindAll(ctx context.Context, query dto.QueryDTO) (res []*response.InfoResponse, page commons.Paginate, err error) {
	infos, total, err := s.infoRepo.FindAll(ctx, query)

	if err != nil {
		return res, page, err
	}

	res = response.ToInfoResponseSlice(infos)

	page = commons.Paginate{
		Page:      query.Page,
		Size:      query.Size,
		TotalPage: total,
	}

	return res, page, nil
}

func (s *infoService) FindPublicInfo(ctx context.Context, query dto.QueryDTO) (res []*response.InfoResponse, page commons.Paginate, err error) {
	infos, total, err := s.infoRepo.FindByStatus(ctx, true, query)

	if err != nil {
		return res, page, err
	}

	res = response.ToInfoResponseSlice(infos)

	page = commons.Paginate{
		Page:      query.Page,
		Size:      query.Size,
		TotalPage: int(total),
	}

	return res, page, nil
}
