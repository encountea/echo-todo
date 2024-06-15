package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	UUID        uuid.UUID `json:"uuid" gorm:"type:uuid;primary_key"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
