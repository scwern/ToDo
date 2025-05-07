package service

import (
	"ToDo/internal/domain/task"
	"fmt"
	"github.com/google/uuid"
)

type TaskRepositoryInterface interface {
	GetAll() []task.Task
	GetById(id uuid.UUID) (*task.Task, error)
	Create(t task.Task) task.Task
	Update(id uuid.UUID, updated task.Task) (*task.Task, error)
	Delete(id uuid.UUID) error
	GetByTitle(title string) (*task.Task, error)
}

type TaskService struct {
	repo TaskRepositoryInterface
}

func NewTaskService(repo TaskRepositoryInterface) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAll() []task.Task {
	return s.repo.GetAll()
}

func (s *TaskService) GetByTitle(title string) (*task.Task, error) {
	return s.repo.GetByTitle(title)
}

func (s *TaskService) GetById(id uuid.UUID) (*task.Task, error) {
	fmt.Println("Fetching task with ID:", id)
	task, err := s.repo.GetById(id)
	if err != nil {
		fmt.Println("Error fetching task:", err)
		return nil, err
	}
	fmt.Println("Found task:", task)
	return task, nil
}

func (s *TaskService) Create(t task.Task) task.Task {
	fmt.Printf("Task received in service: %+v\n", t)
	return s.repo.Create(t)
}

func (s *TaskService) Update(id uuid.UUID, updated task.Task) (*task.Task, error) {
	if updated.Status() == "" {
		updated.SetStatus(task.StatusNew)
	}
	return s.repo.Update(id, updated)
}

func (s *TaskService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
