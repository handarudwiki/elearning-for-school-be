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

type LectureService interface {
	CreateLecture(ctx context.Context, dto dto.CreateLectureDTO) (res response.LectureResponse, err error)
	FindByID(ctx context.Context, id int) (res response.LectureResponse, err error)
	FindAll(ctx context.Context, dto dto.QueryDTO) (res []response.LectureResponse, page commons.Paginate, err error)
	Update(ctx context.Context, dto dto.UpdateLectureDTO, id int) (res response.LectureResponse, err error)
	Delete(ctx context.Context, id int) error
}

type lectureService struct {
	lectureRepo models.LectureRepository
}

func NewLecture(lectureRepo models.LectureRepository) LectureService {
	return &lectureService{
		lectureRepo: lectureRepo,
	}
}

func (s *lectureService) CreateLecture(ctx context.Context, dto dto.CreateLectureDTO) (res response.LectureResponse, err error) {
	lecture := &models.Lecture{
		SubjectID: dto.SubjectID,
		Title:     dto.Title,
		Body:      dto.Body,
		Addition:  dto.Addition,
		Settings:  dto.Settings,
		UserID:    dto.UserID,
	}

	lecture, err = s.lectureRepo.Create(ctx, lecture)
	if err != nil {
		return res, err
	}

	res = response.ToLectureResponse(lecture)

	return res, nil
}

func (s *lectureService) FindByID(ctx context.Context, id int) (res response.LectureResponse, err error) {
	lecture, err := s.lectureRepo.FindByID(ctx, id)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return res, err
	}

	if err != nil {
		return res, commons.ErrNotFound
	}

	res = response.ToLectureResponse(lecture)

	return res, nil
}

func (s *lectureService) FindAll(ctx context.Context, dto dto.QueryDTO) (res []response.LectureResponse, page commons.Paginate, err error) {
	lectures, totalPage, err := s.lectureRepo.FindAll(ctx, dto)
	if err != nil {
		return res, page, err
	}

	page = commons.ToPaginate(dto.Page, dto.Size, totalPage)

	res = response.ToLectureResponseSlice(lectures)

	return res, page, nil
}

func (s *lectureService) Update(ctx context.Context, dto dto.UpdateLectureDTO, id int) (res response.LectureResponse, err error) {
	lecture, err := s.lectureRepo.FindByID(ctx, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return res, err
	}

	if err != nil {
		return res, commons.ErrNotFound
	}

	lecture.SubjectID = dto.SubjectID
	lecture.Title = dto.Title
	lecture.Body = dto.Body
	lecture.Addition = dto.Addition
	lecture.Settings = dto.Settings

	lecture, err = s.lectureRepo.Update(ctx, lecture, id)
	if err != nil {
		return res, err
	}

	res = response.ToLectureResponse(lecture)

	return res, nil
}

func (s *lectureService) Delete(ctx context.Context, id int) error {
	_, err := s.lectureRepo.FindByID(ctx, id)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err != nil {
		return commons.ErrNotFound
	}

	err = s.lectureRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
