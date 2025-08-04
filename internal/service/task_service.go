package service

import (
	"ToDo/internal/domain/task"
	"github.com/google/uuid"
)

// TaskRepositoryInterface определяет поведение хранилища задач.
type TaskRepositoryInterface interface {
	GetAll(userID uuid.UUID) ([]task.Task, error)
	GetById(userID, id uuid.UUID) (*task.Task, error)
	Create(t task.Task) (task.Task, error)
	Update(id uuid.UUID, updated task.Task) (*task.Task, error)
	MarkDeleted(id uuid.UUID) error
	GetByTitle(userID uuid.UUID, title string) (*task.Task, error)
}

// TaskService реализует бизнес-логику работы с задачами.
type TaskService struct {
	Repo TaskRepositoryInterface
}

// NewTaskService создает новый экземпляр TaskService.
func NewTaskService(repo TaskRepositoryInterface) *TaskService {
	return &TaskService{Repo: repo}
}

// GetAll возвращает все задачи пользователя.
func (s *TaskService) GetAll(userID uuid.UUID) ([]task.Task, error) {
	return s.Repo.GetAll(userID)
}

// GetById возвращает задачу по UUID пользователя и задачи.
func (s *TaskService) GetById(userID, id uuid.UUID) (*task.Task, error) {
	return s.Repo.GetById(userID, id)
}

// Create создает новую задачу.
func (s *TaskService) Create(t task.Task) (task.Task, error) {
	return s.Repo.Create(t)
}

// Update обновляет существующую задачу.
func (s *TaskService) Update(id uuid.UUID, updated task.Task) (*task.Task, error) {
	return s.Repo.Update(id, updated)
}

// Delete помечает задачу как удалённую.
func (s *TaskService) Delete(id uuid.UUID) error {
	return s.Repo.MarkDeleted(id)
}

// GetByTitle ищет задачу по названию и UUID пользователя.
func (s *TaskService) GetByTitle(userID uuid.UUID, title string) (*task.Task, error) {
	return s.Repo.GetByTitle(userID, title)
}
