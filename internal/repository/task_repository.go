package repository

import (
	"ToDo/internal/domain/task"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type TaskRepository struct {
	tasks map[uuid.UUID]task.Task
}

var _ TaskRepositoryInterface = (*TaskRepository)(nil)

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[uuid.UUID]task.Task),
	}
}

func (r *TaskRepository) GetAll() []task.Task {
	result := make([]task.Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		result = append(result, t)
	}
	return result
}

func (r *TaskRepository) GetById(id uuid.UUID) (*task.Task, error) {
	t, ok := r.tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	return &t, nil
}

func (r *TaskRepository) Create(t task.Task) task.Task {
	fmt.Printf("Task being saved: %+v\n", t)
	t.SetID(uuid.New())
	r.tasks[t.ID()] = t
	return t
}

func (r *TaskRepository) Update(id uuid.UUID, updated task.Task) (*task.Task, error) {
	if _, ok := r.tasks[id]; !ok {
		return nil, errors.New("task not found")
	}
	updated.SetID(id)
	r.tasks[id] = updated
	return &updated, nil
}

func (r *TaskRepository) Delete(id uuid.UUID) error {
	if _, ok := r.tasks[id]; !ok {
		return errors.New("task not found")
	}
	delete(r.tasks, id)
	return nil
}
