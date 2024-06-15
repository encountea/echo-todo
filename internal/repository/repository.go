package repository

import (
	models "github.com/encountea/echo-todo/internal"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoTask interface {
	Create(task models.Task) error
	GetAll() ([]models.Task, error)
	GetByUUID(id uuid.UUID) (models.Task, error)
	Delete(userId uuid.UUID) error
	Update(task models.Task) error
}

type Repository struct {
	TodoTask
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		TodoTask: NewTodoTaskPostgres(db),
	}
}
