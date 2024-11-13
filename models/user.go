package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/handarudwiki/models/dto"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Details   string    `json:"details"`
	IsActive  bool      `json:"is_active"`
	UID       string    `json:"uid"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.UID = uuid.NewString()
	return nil
}

type UserRepositoy interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUID(ctx context.Context, id int) (*User, error)
	Create(ctx context.Context, user *User) (*User, error)
	FindAll(ctx context.Context, dto dto.QueryDTO) ([]*User, int, error)
	Update(ctx context.Context, id int, user *User) (*User, error)
	Delete(ctx context.Context, id int) error
	FindTeacherByID(ctx context.Context, id int) (*User, error)
	FindStudentByID(ctx context.Context, id int) (*User, error)
}
