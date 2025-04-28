package service

import (
	"ToDo/internal/domain/task"
	"ToDo/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetAll() []task.Taks {
	return s.repo.GetAll()
}

func (s *TaskService) GetById(id int) (*task.Taks, error) {
	return s.repo.GetById(id)
}

func (s *TaskService) Create(t task.Taks) task.Taks {
	return s.repo.Create(t)
}

func (s *TaskService) Update(id int, updated task.Taks) (*task.Taks, error) {
	return s.repo.Update(id, updated)
}

func (s *TaskService) Delete(id int) error {
	return s.repo.Delete(id)
}
