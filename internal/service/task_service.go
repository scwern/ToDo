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
	Repo TaskRepositoryInterface
}

func NewTaskService(repo TaskRepositoryInterface) *TaskService {
	return &TaskService{Repo: repo}
}

func (s *TaskService) GetAll(userID uuid.UUID) ([]task.Task, error) {
	return s.Repo.GetAll(userID)
}

func (s *TaskService) GetById(userID, id uuid.UUID) (*task.Task, error) {
	return s.Repo.GetById(userID, id)
}

func (s *TaskService) Create(t task.Task) (task.Task, error) {
	return s.Repo.Create(t)
}

func (s *TaskService) Update(id uuid.UUID, updated task.Task) (*task.Task, error) {
	return s.Repo.Update(id, updated)
}

func (s *TaskService) Delete(id uuid.UUID) error {
	return s.Repo.MarkDeleted(id)
}

func (s *TaskService) GetByTitle(userID uuid.UUID, title string) (*task.Task, error) {
	return s.Repo.GetByTitle(userID, title)
}
