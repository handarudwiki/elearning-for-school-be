package repositories

import (
	"context"
	"fmt"
	"math"

	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/models"
	"github.com/handarudwiki/models/commons"
	"github.com/handarudwiki/models/dto"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) models.UserRepositoy {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByUID(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	fmt.Println(user)

	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, id int, user *models.User) (*models.User, error) {
	err := r.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Update("is_active", user.IsActive).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	err := r.db.Where("id = ?", id).Delete(&models.User{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindAll(ctx context.Context, dto dto.QueryDTO) ([]*models.User, int, error) {
	var users []*models.User

	err := r.db.Scopes(helpers.Paginate(dto.Page, dto.Size), helpers.Search(*dto.Search), helpers.FilterIsActive(dto.Is_active), helpers.FilterRole(dto.Role)).Find(&users).Error

	var totalUsers int64
	err = r.db.Model(users).Scopes(helpers.Search(*dto.Search), helpers.FilterIsActive(dto.Is_active), helpers.FilterRole(dto.Role)).Count(&totalUsers).Error

	if err != nil {
		return nil, 0, err
	}

	totalPage := math.Ceil(float64(totalUsers) / float64(dto.Size))
	if err != nil {
		return nil, int(totalPage), err
	}

	return users, int(totalPage), nil
}

func (r *UserRepository) FindTeacherByID(ctx context.Context, id int) (*models.User, error) {
	var users *models.User
	err := r.db.Where("id = ?", id).Where("role = ?", commons.ROLETEACHER).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) FindStudentByID(ctx context.Context, id int) (*models.User, error) {
	var users *models.User
	err := r.db.Where("id = ?", id).Where("role = ?", commons.ROLESTUDENT).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}
