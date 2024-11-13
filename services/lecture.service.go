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
