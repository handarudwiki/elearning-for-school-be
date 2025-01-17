package services

import (
	"context"

	"github.com/handarudwiki/models"
	"github.com/handarudwiki/models/commons"
	"github.com/handarudwiki/models/dto"
	"github.com/handarudwiki/models/response"
)

type LectureCommentService interface {
	Create(ctx context.Context, comment *dto.CreateLectureCommentDTO) (*response.LectureCommentResponse, error)
	FindAll(ctx context.Context, lectureID int, query dto.QueryDTO) ([]response.LectureCommentResponse, commons.Paginate, error)
}

type lectureCommentService struct {
	lectureCommentRepo models.LectureCommentRepository
	lectureService     LectureService
}

func NewLectureComment(lectureCommentRepo models.LectureCommentRepository, lectureService LectureService) LectureCommentService {
	return &lectureCommentService{
		lectureCommentRepo: lectureCommentRepo,
		lectureService:     lectureService,
	}
}

func (s *lectureCommentService) Create(ctx context.Context, comment *dto.CreateLectureCommentDTO) (*response.LectureCommentResponse, error) {

	_, err := s.lectureService.FindByID(ctx, int(comment.LectureID))

	if err != nil {
		return nil, err
	}

	lectureComment := &models.LectureComment{
		LectureID: comment.LectureID,
		Content:   comment.Content,
		UserID:    comment.UserID,
	}

	lectureComment, err = s.lectureCommentRepo.Create(ctx, lectureComment)
	if err != nil {
		return nil, err
	}

	res := response.ToLectureCommentResponse(lectureComment)

	return &res, nil
}

func (s *lectureCommentService) FindAll(ctx context.Context, lectureID int, query dto.QueryDTO) ([]response.LectureCommentResponse, commons.Paginate, error) {
	comments, count, err := s.lectureCommentRepo.FindByLectureID(ctx, lectureID, query)
	if err != nil {
		return nil, commons.Paginate{}, err
	}

	var res []response.LectureCommentResponse
	for _, comment := range comments {
		res = append(res, response.ToLectureCommentResponse(comment))
	}

	paginate := commons.ToPaginate(query.Page, query.Size, int(count))
	return res, paginate, nil
}
