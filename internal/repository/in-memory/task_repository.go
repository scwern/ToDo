package in_memory

import (
	"ToDo/internal/domain/task"
	"errors"
	"github.com/google/uuid"
)

type TaskRepository struct {
	tasks map[uuid.UUID]task.Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[uuid.UUID]task.Task),
	}
}

func (r *TaskRepository) GetByTitle(userID uuid.UUID, title string) (*task.Task, error) {
	for _, t := range r.tasks {
		if t.UserID() == userID && t.Title() == title {
			return &t, nil
		}
	}
	return nil, errors.New("task with this title not found")
}

func (r *TaskRepository) GetAll(userID uuid.UUID) ([]task.Task, error) {
	result := make([]task.Task, 0)
	for _, t := range r.tasks {
		if t.UserID() == userID {
			result = append(result, t)
		}
	}
	return result, nil
}

func (r *TaskRepository) GetById(userID, id uuid.UUID) (*task.Task, error) {
	t, ok := r.tasks[id]
	if !ok || t.UserID() != userID {
		return nil, errors.New("task not found")
	}
	return &t, nil
}

func (r *TaskRepository) Create(t task.Task) (task.Task, error) {
	t.SetID(uuid.New())
	r.tasks[t.ID()] = t
	return t, nil
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
