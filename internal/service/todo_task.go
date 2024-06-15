package service

import (
	"github.com/encountea/echo-todo/internal/repository"
	"github.com/encountea/echo-todo/internal"
	"github.com/google/uuid"
)

type TodoTaskService struct {
	repo repository.TodoTask
}

func NewTodoTaskService(repo repository.TodoTask) *TodoTaskService {
	return &TodoTaskService{repo: repo}
}

func (s *TodoTaskService) GetAll() ([]models.Task, error) {
	return s.repo.GetAll()
}

func (s *TodoTaskService) GetByUUID(id uuid.UUID) (models.Task, error) {
	return s.repo.GetByUUID(id)
}

func (s *TodoTaskService) Create(task models.Task) error {
	return s.repo.Create(task)
}

func (s *TodoTaskService) Update(task models.Task) error {
	return s.repo.Update(task)
}

func (s *TodoTaskService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
