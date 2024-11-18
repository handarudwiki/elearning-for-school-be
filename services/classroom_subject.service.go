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

type ClassroomSubjectService interface {
	FindByTeacherID(ctx context.Context, teacherID int) (res []response.ClassroomSubjectResponse, err error)
	Create(ctx context.Context, dto dto.CreateClassrooomSubject) (res response.ClassroomSubjectResponse, err error)
	Delete(ctx context.Context, id int) error
}

type classroomSubjectService struct {
	classroomSubjectRepo models.ClassroomSubjectRepository
	clssroomService      ClassroomService
	subjectService       SubjectService
}

func NewClassroomSubject(classroomSubjectRepo models.ClassroomSubjectRepository, clssroomService ClassroomService, subjectService SubjectService) ClassroomSubjectService {
	return &classroomSubjectService{
		classroomSubjectRepo: classroomSubjectRepo,
		clssroomService:      clssroomService,
		subjectService:       subjectService,
	}
}

func (s *classroomSubjectService) FindByTeacherID(ctx context.Context, teacherID int) (res []response.ClassroomSubjectResponse, err error) {
	classroomSubjects, err := s.classroomSubjectRepo.FindByTeacherID(ctx, teacherID)
	if err != nil {
		return res, err
	}

	if len(classroomSubjects) == 0 {
		return res, nil
	}

	res = response.ToclassroomSubjectResponseSlice(classroomSubjects)

	return res, nil
}

func (s *classroomSubjectService) Create(ctx context.Context, dto dto.CreateClassrooomSubject) (res response.ClassroomSubjectResponse, err error) {
	_, err = s.clssroomService.FindByID(ctx, int(dto.ClassroomID))
	if err != nil {
		return res, err
	}

	_, err = s.subjectService.FindByID(ctx, int(dto.SubjectID))
	if err != nil {
		return res, err
	}

	classroomSubject := &models.ClassroomSubject{
		ClassroomID: dto.ClassroomID,
		SubjectID:   dto.SubjectID,
		TeacherID:   dto.TeacherID,
	}

	newClassroomSubject, err := s.classroomSubjectRepo.Create(ctx, classroomSubject)
	if err != nil {
		return res, err
	}

	res = response.ToClassroomSubjectResponse(newClassroomSubject)

	return res, nil
}

func (s *classroomSubjectService) Delete(ctx context.Context, id int) error {

	_, err := s.classroomSubjectRepo.FindByID(ctx, id)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err != nil {
		return commons.ErrNotFound
	}

	err = s.classroomSubjectRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
