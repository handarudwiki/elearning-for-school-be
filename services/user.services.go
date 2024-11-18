package services

import (
	"context"
	"errors"

	"github.com/handarudwiki/models"
	"github.com/handarudwiki/models/commons"
	"github.com/handarudwiki/models/dto"
	"github.com/handarudwiki/models/response"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	Login(ctx context.Context, loginDto dto.LoginDTO) (response.LoginResponse, error)
	Me(ctx context.Context, id int) (response.UserResponse, error)
	CreateTeacher(ctx context.Context, dto dto.CreateUserDTO) (res response.UserResponse, err error)
	CreateStudent(ctx context.Context, dto dto.CreateUserDTO) (res response.UserResponse, err error)
	GetAllTeacher(ctx context.Context, dto dto.QueryDTO) (res []response.UserResponse, page commons.Paginate, err error)
	GetAllStudent(ctx context.Context, dto dto.QueryDTO) (res []response.UserResponse, page commons.Paginate, err error)
	GetJwtService() JWTService
	GetUser(ctx context.Context, id int) (res response.UserResponse, err error)
}

type userService struct {
	userRepo   models.UserRepositoy
	jwtService JWTService
}

func NewUser(userRepo models.UserRepositoy, jwtService JWTService) UserService {
	return &userService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (s *userService) Login(ctx context.Context, loginDto dto.LoginDTO) (td response.LoginResponse, err error) {
	user, err := s.userRepo.FindByEmail(ctx, loginDto.Email)

	if err != nil {
		return response.LoginResponse{}, commons.ErrCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password))

	if err != nil {
		return response.LoginResponse{}, commons.ErrCredentials
	}

	accessToken, err := s.jwtService.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return response.LoginResponse{}, err
	}

	refreshToken, err := s.jwtService.GenrateRefreshToken(user.ID, user.Role)
	if err != nil {
		return response.LoginResponse{}, err
	}

	td = response.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return td, nil
}

func (s *userService) Me(ctx context.Context, id int) (res response.UserResponse, err error) {
	user, err := s.userRepo.FindByUID(ctx, id)
	if err != nil {
		return
	}

	res = response.ToUserResponse(user)

	return
}

func (s *userService) CreateTeacher(ctx context.Context, dto dto.CreateUserDTO) (res response.UserResponse, err error) {
	newUser := &models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
		Details:  dto.Details,
		IsActive: true,
		Role:     commons.ROLETEACHER,
	}

	_, err = s.userRepo.FindByEmail(ctx, newUser.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	if err == nil {
		err = commons.ErrConflict
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	newUser.Password = string(hashedPassword)

	newUser, err = s.userRepo.Create(ctx, newUser)
	if err != nil {
		return
	}
	res = response.ToUserResponse(newUser)

	return
}

func (s *userService) CreateStudent(ctx context.Context, dto dto.CreateUserDTO) (res response.UserResponse, err error) {
	newUser := &models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
		Details:  dto.Details,
		IsActive: true,
		Role:     commons.ROLESTUDENT,
	}

	_, err = s.userRepo.FindByEmail(ctx, newUser.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	if err == nil {
		err = commons.ErrConflict
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	newUser.Password = string(hashedPassword)

	newUser, err = s.userRepo.Create(ctx, newUser)
	if err != nil {
		return
	}
	res = response.ToUserResponse(newUser)

	return
}

func (s *userService) GetAllTeacher(ctx context.Context, dto dto.QueryDTO) (res []response.UserResponse, page commons.Paginate, err error) {
	teacher := commons.ROLETEACHER
	dto.Role = &teacher
	teachers, totalPage, err := s.userRepo.FindAll(ctx, dto)
	if err != nil {
		return
	}

	page = commons.ToPaginate(dto.Page, dto.Size, totalPage)

	res = response.ToUserResponseSlice(teachers)

	return
}

func (s *userService) GetAllStudent(ctx context.Context, dto dto.QueryDTO) (res []response.UserResponse, page commons.Paginate, err error) {
	student := commons.ROLESTUDENT
	dto.Role = &student
	students, totalPage, err := s.userRepo.FindAll(ctx, dto)
	if err != nil {
		return
	}

	page = commons.ToPaginate(dto.Page, dto.Size, totalPage)

	res = response.ToUserResponseSlice(students)

	return
}

func (s *userService) GetUser(ctx context.Context, id int) (res response.UserResponse, err error) {
	user, err := s.userRepo.FindByUID(ctx, id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	if err != nil {
		err = commons.ErrNotFound
		return
	}

	res = response.ToUserResponse(user)

	return

}

func (s *userService) GetJwtService() JWTService {
	return s.jwtService
}
