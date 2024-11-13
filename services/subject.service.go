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

type SubjectService interface {
	FindByID(ctx context.Context, id int) (res response.SubjectResponse, err error)
	Create(dto *dto.CreateSubjectDTO) (res response.SubjectResponse, err error)
	FindAll(ctx context.Context, dto dto.QueryDTO) (res []response.SubjectResponse, page commons.Paginate, err error)
	Update(ctx context.Context, id int, dto dto.UpdateSubjectDTO) (res response.SubjectResponse, err error)
	Delete(ctx context.Context, id int) error
}

type subjectService struct {
	subjectRepo models.SubjectRepository
}

func NewSubject(subjectRepo models.SubjectRepository) SubjectService {
	return &subjectService{
		subjectRepo: subjectRepo,
	}
}

func (s *subjectService) FindByID(ctx context.Context, id int) (res response.SubjectResponse, err error) {
	subject, err := s.subjectRepo.FindByID(ctx, id)
	if err != nil {
		return res, err
	}

	return response.ToSubjectResponse(subject), nil
}

func (s *subjectService) Create(dto *dto.CreateSubjectDTO) (res response.SubjectResponse, err error) {
	subject := &models.Subject{
		Name:        dto.Name,
		Description: dto.Description,
	}

	subject, err = s.subjectRepo.Create(subject)
	if err != nil {
		return res, err
	}

	return response.ToSubjectResponse(subject), nil
}

func (s *subjectService) FindAll(ctx context.Context, dto dto.QueryDTO) (res []response.SubjectResponse, page commons.Paginate, err error) {
	subjects, totalPage, err := s.subjectRepo.FindAll(ctx, dto)
	if err != nil {
		return nil, page, err
	}

	page = commons.ToPaginate(dto.Page, dto.Size, int(totalPage))

	res = response.TosSubjectResponseSlice(subjects)

	return
}

func (s *subjectService) Update(ctx context.Context, id int, dto dto.UpdateSubjectDTO) (res response.SubjectResponse, err error) {
	subject, err := s.subjectRepo.FindByID(ctx, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return res, err
	}

	if err != nil {
		return res, commons.ErrNotFound
	}

	subject.Name = dto.Name
	subject.Description = dto.Description

	subject, err = s.subjectRepo.Update(ctx, id, subject)
	if err != nil {
		return res, err
	}

	return response.ToSubjectResponse(subject), nil
}

func (s *subjectService) Delete(ctx context.Context, id int) error {
	_, err := s.subjectRepo.FindByID(ctx, id)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err != nil {
		return commons.ErrNotFound
	}

	err = s.subjectRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
