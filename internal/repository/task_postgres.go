package repository

import (
	models "github.com/encountea/echo-todo/internal"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoTaskPostgres struct {
	db *gorm.DB
}

func NewTodoTaskPostgres(db *gorm.DB) *TodoTaskPostgres {
	return &TodoTaskPostgres{db: db}
}

func (r *TodoTaskPostgres) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TodoTaskPostgres) GetByUUID(id uuid.UUID) (models.Task, error) {
	var task models.Task
	if err := r.db.First(&task, "uuid = ?", id).Error; err != nil {
		return task, err
	}
	return task, nil
}

func (r *TodoTaskPostgres) Create(task models.Task) error {
	return r.db.Create(&task).Error
}

func (r *TodoTaskPostgres) Update(task models.Task) error {
	return r.db.Save(&task).Error
}

func (r *TodoTaskPostgres) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Task{}, "uuid = ?", id).Error
}
