package models

import (
	"context"
	"time"

	"github.com/handarudwiki/models/dto"
)

type Classroom struct {
	ID        uint      `json:"id"`
	TeacherID uint      `json:"teacher_id"`
	Teacher   *User     `json:"teacher"`
	Name      string    `json:"name"`
	Grade     string    `json:"grade"`
	Group     string    `json:"group"`
	Settings  string    `json:"settings"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ClassroomRepository interface {
	FindByID(ctx context.Context, id int) (*Classroom, error)
	Create(ctx context.Context, classroom *Classroom) (*Classroom, error)
	Update(ctx context.Context, classroom *Classroom, id int) (*Classroom, error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context, dto dto.QueryDTO) ([]*Classroom, int64, error)
}
