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

type ClassroomService interface {
	FindByID(ctx context.Context, id int) (res response.ClassroomResponse, err error)
	Create(ctx context.Context, dto dto.CreateClassroomDTO) (res response.ClassroomResponse, err error)
	Update(ctx context.Context, dto dto.UpdateClassroomDTO, id int) (res response.ClassroomResponse, err error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context, dto dto.QueryDTO) (res []response.ClassroomResponse, page commons.Paginate, err error)
	FindClassroomStudent(ctx context.Context, classroomID int) (res []response.ClassroomStudentResponse, err error)
	AssignStudentClassroom(ctx context.Context, dto dto.CreateClassroomStudentDTO) (res response.ClassroomStudentResponse, err error)
	FindByTeacherID(ctx context.Context, teacherID int) (res []response.ClassroomResponse, err error)
}

type classroomService struct {
	classroomRepo        models.ClassroomRepository
	userRepo             models.UserRepositoy
	classroomStudentRepo models.ClassroomStudentRepository
}

func NewClassroom(classroomRepo models.ClassroomRepository, userRepo models.UserRepositoy, classroomStudentRepo models.ClassroomStudentRepository) ClassroomService {
	return &classroomService{
		classroomRepo:        classroomRepo,
		userRepo:             userRepo,
		classroomStudentRepo: classroomStudentRepo,
	}
}

func (s *classroomService) FindByID(ctx context.Context, id int) (res response.ClassroomResponse, err error) {
	classroom, err := s.classroomRepo.FindByID(ctx, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return res, err
	}

	if err != nil {
		return res, commons.ErrNotFound
	}

	res = response.ToClassroomResponse(classroom)

	return res, nil
}

func (s *classroomService) Create(ctx context.Context, dto dto.CreateClassroomDTO) (res response.ClassroomResponse, err error) {
	user, err := s.userRepo.FindTeacherByID(ctx, int(dto.TeacherID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return res, err
	}
	if err != nil {
		return res, commons.ErrNotFound
	}

	if user.ID == 0 {
		return res, commons.ErrNotFound
	}

	classroom := &models.Classroom{
		Name:      dto.Name,
		TeacherID: dto.TeacherID,
		Grade:     dto.Grade,
		Group:     dto.Group,
		Settings:  dto.Settings,
	}

	classroom, err = s.classroomRepo.Create(ctx, classroom)
	if err != nil {
		return res, err
	}

	res = response.ToClassroomResponse(classroom)

	return res, nil
}

func (s *classroomService) Update(ctx context.Context, dto dto.UpdateClassroomDTO, id int) (res response.ClassroomResponse, err error) {
	classroom, err := s.classroomRepo.FindByID(ctx, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	if err != nil {
		return res, commons.ErrNotFound
	}

	teacher, err := s.userRepo.FindTeacherByID(ctx, int(dto.TeacherID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return res, err
	}

	if err != nil {
		return res, commons.ErrNotFound
	}

	if teacher.ID == 0 {
		return res, commons.ErrNotFound
	}

	classroom.TeacherID = dto.TeacherID
	classroom.Grade = dto.Grade
	classroom.Group = dto.Group
	classroom.Settings = dto.Settings

	classroom, err = s.classroomRepo.Update(ctx, classroom, id)

	if err != nil {
		return res, err
	}

	res = response.ToClassroomResponse(classroom)

	return res, nil
}

func (s *classroomService) Delete(ctx context.Context, id int) error {
	_, err := s.classroomRepo.FindByID(ctx, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err != nil {
		return commons.ErrNotFound
	}

	err = s.classroomRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *classroomService) FindAll(ctx context.Context, dto dto.QueryDTO) (res []response.ClassroomResponse, page commons.Paginate, err error) {
	classrooms, totalPage, err := s.classroomRepo.FindAll(ctx, dto)
	if err != nil {
		return res, page, err
	}

	page = commons.ToPaginate(dto.Page, dto.Size, int(totalPage))

	res = response.ToClassroomResponseSlice(classrooms)

	return res, page, nil
}

func (s *classroomService) FindClassroomStudent(ctx context.Context, classroomID int) (res []response.ClassroomStudentResponse, err error) {
	classrooms, err := s.classroomStudentRepo.FindByClassroomID(ctx, classroomID)
	if err != nil {
		return res, err
	}

	res = response.ToClassroomStudentResponseSlice(classrooms)

	return res, nil
}

func (s *classroomService) AssignStudentClassroom(ctx context.Context, dto dto.CreateClassroomStudentDTO) (res response.ClassroomStudentResponse, err error) {
	_, err = s.classroomRepo.FindByID(ctx, dto.ClassroomID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	if err != nil {
		err = commons.ErrNotFound
		return
	}

	_, err = s.userRepo.FindByUID(ctx, dto.StudentID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	if err != nil {
		err = commons.ErrNotFound
		return
	}

	classroomStudent := &models.ClassroomStudent{
		ClassroomID: uint(dto.ClassroomID),
		StudentID:   uint(dto.StudentID),
	}

	_, err = s.classroomStudentRepo.Create(ctx, classroomStudent)
	if err != nil {
		return
	}

	res = response.ToClassroomStudentResponse(classroomStudent)

	return
}

func (s *classroomService) FindByTeacherID(ctx context.Context, teacherID int) (res []response.ClassroomResponse, err error) {
	classrooms, err := s.classroomRepo.FindByTeacherID(ctx, teacherID)
	if err != nil {
		return res, err
	}

	res = response.ToClassroomResponseSlice(classrooms)

	return res, nil
}
