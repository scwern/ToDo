package service

import (
	"ToDo/internal/domain/task"
	"github.com/google/uuid"
)

type TaskRepositoryInterface interface {
	GetAll(userID uuid.UUID) ([]task.Task, error)
	GetById(userID, id uuid.UUID) (*task.Task, error)
	Create(t task.Task) (task.Task, error)
	Update(id uuid.UUID, updated task.Task) (*task.Task, error)
	MarkDeleted(id uuid.UUID) error
	GetByTitle(userID uuid.UUID, title string) (*task.Task, error)
}

type TaskService struct {
	repo TaskRepositoryInterface
}

func NewTaskService(repo TaskRepositoryInterface) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAll(userID uuid.UUID) ([]task.Task, error) {
	return s.repo.GetAll(userID)
}

func (s *TaskService) GetById(userID, id uuid.UUID) (*task.Task, error) {
	return s.repo.GetById(userID, id)
}

func (s *TaskService) Create(t task.Task) (task.Task, error) {
	return s.repo.Create(t)
}

func (s *TaskService) Update(id uuid.UUID, updated task.Task) (*task.Task, error) {
	return s.repo.Update(id, updated)
}

func (s *TaskService) Delete(id uuid.UUID) error {
	return s.repo.MarkDeleted(id)
}

func (s *TaskService) GetByTitle(userID uuid.UUID, title string) (*task.Task, error) {
	return s.repo.GetByTitle(userID, title)
}
