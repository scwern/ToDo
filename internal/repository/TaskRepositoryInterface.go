package repository

import (
	"ToDo/internal/domain/task"
	"github.com/google/uuid"
)

type TaskRepositoryInterface interface {
	GetAll() []task.Task
	GetById(id uuid.UUID) (*task.Task, error)
	Create(t task.Task) task.Task
	Update(id uuid.UUID, updated task.Task) (*task.Task, error)
	Delete(id uuid.UUID) error
}
