package repositories

import (
	"context"

	"github.com/handarudwiki/helpers"
	"github.com/handarudwiki/models"
	"github.com/handarudwiki/models/dto"
	"gorm.io/gorm"
)

type subjectRepository struct {
	db *gorm.DB
}

func NewSubject(db *gorm.DB) models.SubjectRepository {
	return &subjectRepository{
		db: db,
	}
}

func (r *subjectRepository) FindByID(ctx context.Context, id int) (res *models.Subject, err error) {
	var subject models.Subject
	err = r.db.Where("id = ?", id).First(&subject).Error

	if err != nil {
		return nil, err
	}

	return &subject, nil
}

func (r *subjectRepository) Create(subject *models.Subject) (*models.Subject, error) {
	err := r.db.Create(subject).Error
	if err != nil {
		return nil, err
	}

	return subject, nil
}

func (r *subjectRepository) FindAll(ctx context.Context, dto dto.QueryDTO) (subjects []*models.Subject, totalPage int64, err error) {
	var subject []*models.Subject

	err = r.db.Scopes(helpers.Paginate(dto.Page, dto.Size), helpers.Search(*dto.Search)).Find(&subject).Error

	if err != nil {
		return nil, totalPage, err
	}

	err = r.db.Model(&models.Subject{}).Scopes(helpers.Search(*dto.Search)).Count(&totalPage).Error

	return subject, totalPage, nil
}

func (r *subjectRepository) Update(ctx context.Context, id int, subject *models.Subject) (*models.Subject, error) {
	err := r.db.Where("id = ?", id).Updates(subject).Error
	if err != nil {
		return nil, err
	}

	return subject, nil
}

func (r *subjectRepository) Delete(ctx context.Context, id int) error {
	err := r.db.Where("id = ?", id).Delete(&models.Subject{}).Error
	if err != nil {
		return err
	}

	return nil
}
