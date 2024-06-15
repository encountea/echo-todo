package service

import (
	models "github.com/encountea/echo-todo/internal"
	"github.com/encountea/echo-todo/internal/repository"
	"github.com/google/uuid"
)

type TodoTask interface {
	GetAll() ([]models.Task, error)
	GetByUUID(id uuid.UUID) (models.Task, error)
	Create(todo models.Task) error
	Update(todo models.Task) error
	Delete(id uuid.UUID) error
}

type Service struct {
	TodoTask
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TodoTask: NewTodoTaskService(repos.TodoTask),
	}
}
