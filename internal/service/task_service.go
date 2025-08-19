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

// NewTaskService creates a new task service instance
// @Summary Create task service
// @Tags Tasks
// @Router /task-service [post]
func NewTaskService(repo TaskRepositoryInterface) *TaskService {
	return &TaskService{Repo: repo}
}

// GetAll returns all user tasks
// @Summary Get all tasks
// @Tags Tasks
// @Param userID path string true "User UUID"
// @Success 200 {array} task.Task
// @Router /tasks/{userID} [get]
func (s *TaskService) GetAll(userID uuid.UUID) ([]task.Task, error) {
	return s.Repo.GetAll(userID)
}

// GetById returns task by ID
// @Summary Get task by ID
// @Tags Tasks
// @Param userID path string true "User UUID"
// @Param id path string true "Task UUID"
// @Success 200 {object} task.Task
// @Router /tasks/{userID}/{id} [get]
func (s *TaskService) GetById(userID, id uuid.UUID) (*task.Task, error) {
	return s.Repo.GetById(userID, id)
}

// Create creates new task
// @Summary Create task
// @Tags Tasks
// @Param input body task.Task true "Task data"
// @Success 200 {object} task.Task
// @Router /tasks [post]
func (s *TaskService) Create(t task.Task) (task.Task, error) {
	return s.Repo.Create(t)
}

// Update updates existing task
// @Summary Update task
// @Tags Tasks
// @Param id path string true "Task UUID"
// @Param input body task.Task true "Updated task data"
// @Success 200 {object} task.Task
// @Router /tasks/{id} [put]
func (s *TaskService) Update(id uuid.UUID, updated task.Task) (*task.Task, error) {
	return s.Repo.Update(id, updated)
}

// Delete marks task as deleted
// @Summary Delete task
// @Tags Tasks
// @Param id path string true "Task UUID"
// @Success 200
// @Router /tasks/{id} [delete]
func (s *TaskService) Delete(id uuid.UUID) error {
	return s.Repo.MarkDeleted(id)
}

// GetByTitle finds task by title
// @Summary Find task by title
// @Tags Tasks
// @Param userID path string true "User UUID"
// @Param title query string true "Task title"
// @Success 200 {object} task.Task
// @Router /tasks/{userID}/search [get]
func (s *TaskService) GetByTitle(userID uuid.UUID, title string) (*task.Task, error) {
	return s.Repo.GetByTitle(userID, title)
}
