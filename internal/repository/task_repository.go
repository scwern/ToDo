package repository

import (
	"ToDo/internal/domain/task"
	"errors"
)

type TaskRepository struct {
	tasks  []task.Task
	nextID int
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks:  make([]task.Task, 0),
		nextID: 1,
	}
}
func (r *TaskRepository) GetAll() []task.Task {
	return r.tasks
}

func (r *TaskRepository) GetById(id int) (*task.Task, error) {
	for _, t := range r.tasks {
		if t.ID == id {
			return &t, nil
		}
	}
	return nil, errors.New("task not found")
}

func (r *TaskRepository) Create(t task.Task) task.Task {
	if t.Status == "" {
		t.Status = task.StatusNew
	}

	t.ID = r.nextID
	r.nextID++
	r.tasks = append(r.tasks, t)
	return t
}

func (r *TaskRepository) Update(id int, updated task.Task) (*task.Task, error) {
	for i, t := range r.tasks {
		if t.ID == id {
			if updated.Status == "" {
				updated.Status = t.Status
			}
			updated.ID = id
			r.tasks[i] = updated
			return &r.tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

func (r *TaskRepository) Delete(id int) error {
	for i, t := range r.tasks {
		if t.ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
