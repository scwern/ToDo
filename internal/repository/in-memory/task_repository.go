package in_memory

import (
	"ToDo/internal/domain/task"
	"errors"
	"github.com/google/uuid"
	"log"
	"sync"
)

type TaskRepository struct {
	tasks    map[uuid.UUID]task.Task
	deleteCh chan uuid.UUID
	wg       sync.WaitGroup
	once     sync.Once
	stopCh   chan struct{}
	mu       sync.Mutex
}

func NewTaskRepository() *TaskRepository {
	repo := &TaskRepository{
		tasks:    make(map[uuid.UUID]task.Task),
		deleteCh: make(chan uuid.UUID, 10),
		stopCh:   make(chan struct{}),
	}
	repo.wg.Add(1)
	go repo.processDeletions()
	return repo
}

func (r *TaskRepository) processDeletions() {
	defer r.wg.Done()
	var batch []uuid.UUID
	for {
		select {
		case id, ok := <-r.deleteCh:
			if !ok {
				if len(batch) > 0 {
					r.deleteBatch(batch)
				}
				return
			}
			batch = append(batch, id)
			if len(batch) >= 10 {
				r.deleteBatch(batch)
				batch = nil
			}
		case <-r.stopCh:
			if len(batch) > 0 {
				r.deleteBatch(batch)
			}
			return
		}
	}
}

func (r *TaskRepository) deleteBatch(ids []uuid.UUID) {
	log.Printf("Deleting batch of %d tasks: %v", len(ids), ids)

	r.mu.Lock()
	defer r.mu.Unlock()
	for _, id := range ids {
		delete(r.tasks, id)
	}
}

func (r *TaskRepository) Close() {
	r.once.Do(func() {
		close(r.stopCh)
		r.wg.Wait()
	})
}

func (r *TaskRepository) MarkDeleted(id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	t, ok := r.tasks[id]
	if !ok {
		return errors.New("task not found")
	}

	t.SetDeleted(true)
	r.tasks[id] = t
	log.Printf("Task %v marked as deleted, sending to channel", id)

	select {
	case r.deleteCh <- id:
	default:
		go func() {
			r.deleteCh <- id
		}()
	}
	return nil
}

func (r *TaskRepository) GetByTitle(userID uuid.UUID, title string) (*task.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, t := range r.tasks {
		if t.UserID() == userID && t.Title() == title && !t.IsDeleted() {
			return &t, nil
		}
	}
	return nil, errors.New("task not found")
}

func (r *TaskRepository) GetAll(userID uuid.UUID) ([]task.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	result := make([]task.Task, 0)
	for _, t := range r.tasks {
		if t.UserID() == userID && !t.IsDeleted() {
			result = append(result, t)
		}
	}
	return result, nil
}

func (r *TaskRepository) GetById(userID, id uuid.UUID) (*task.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	t, ok := r.tasks[id]
	if !ok || t.UserID() != userID || t.IsDeleted() {
		return nil, errors.New("task not found")
	}
	return &t, nil
}

func (r *TaskRepository) Create(t task.Task) (task.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	t.SetID(uuid.New())
	r.tasks[t.ID()] = t
	return t, nil
}

func (r *TaskRepository) Update(id uuid.UUID, updated task.Task) (*task.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.tasks[id]; !ok {
		return nil, errors.New("task not found")
	}
	updated.SetID(id)
	r.tasks[id] = updated
	return &updated, nil
}

func (r *TaskRepository) Delete(id uuid.UUID) error {
	return r.MarkDeleted(id)
}
