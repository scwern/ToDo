package dbstorage

import (
	"ToDo/internal/domain/task"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log"
	"sync"
)

type TaskRepository struct {
	db       *pgx.Conn
	deleteCh chan uuid.UUID
	wg       sync.WaitGroup
	once     sync.Once
	stopCh   chan struct{}
}

func NewTaskRepository(db *pgx.Conn) *TaskRepository {
	repo := &TaskRepository{
		db:       db,
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
	batchSize := 10

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
			if len(batch) >= batchSize {
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

func (r *TaskRepository) deleteBatch(ids []uuid.UUID) error {
	log.Printf("Starting batch delete of %d tasks: %v", len(ids), ids)

	ctx := context.Background()
	tx, err := r.db.Begin(ctx)
	if err != nil {
		log.Printf("Failed to begin transaction: %v", err)
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `DELETE FROM tasks WHERE uid = ANY($1)`, ids)
	if err != nil {
		log.Printf("Failed to delete tasks: %v", err)
		return fmt.Errorf("failed to delete tasks: %w", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Printf("Failed to commit transaction: %v", err)
	} else {
		log.Printf("Successfully deleted %d tasks", len(ids))
	}
	return err
}

func (r *TaskRepository) Close() {
	r.once.Do(func() {
		close(r.stopCh)
		r.wg.Wait()
		log.Println("Task repository closed")
	})
}

func (r *TaskRepository) MarkDeleted(id uuid.UUID) error {
	ctx := context.Background()
	_, err := r.db.Exec(ctx, `UPDATE tasks SET deleted = true WHERE uid = $1`, id)
	if err != nil {
		log.Printf("Failed to mark task %v as deleted: %v", id, err)
		return fmt.Errorf("failed to mark task as deleted: %w", err)
	}

	log.Printf("Task %v marked as deleted, sending to channel", id)
	select {
	case r.deleteCh <- id:
		log.Printf("Task %v added to delete channel", id)
	default:
		go func() {
			log.Printf("Channel full, retrying to add task %v", id)
			r.deleteCh <- id
		}()
	}
	return nil
}

func (r *TaskRepository) Delete(id uuid.UUID) error {
	return r.MarkDeleted(id)
}

func (r *TaskRepository) Create(t task.Task) (task.Task, error) {
	query := `INSERT INTO tasks (uid, user_id, title, description, status) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(context.Background(), query,
		t.ID(), t.UserID(), t.Title(), t.Description(), int(t.Status()))
	if err != nil {
		log.Printf("Failed to create task: %v", err)
		return task.Task{}, fmt.Errorf("failed to insert task: %w", err)
	}
	log.Printf("Task created: %v", t.ID())
	return t, nil
}

func (r *TaskRepository) GetByTitle(userID uuid.UUID, title string) (*task.Task, error) {
	query := `SELECT uid, user_id, title, description, status FROM tasks WHERE user_id = $1 AND title = $2 AND deleted = false LIMIT 1`
	row := r.db.QueryRow(context.Background(), query, userID, title)

	var tid, uidUserID uuid.UUID
	var titleDB, description string
	var status int
	err := row.Scan(&tid, &uidUserID, &titleDB, &description, &status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Printf("Task not found by title: %s", title)
			return nil, fmt.Errorf("task not found: %w", err)
		}
		log.Printf("Error getting task by title: %v", err)
		return nil, fmt.Errorf("failed to get task by title: %w", err)
	}

	t := task.NewTask(titleDB, description, task.Status(status))
	t.SetID(tid)
	t.SetUserID(uidUserID)

	return &t, nil
}

func (r *TaskRepository) GetAll(userID uuid.UUID) ([]task.Task, error) {
	query := `SELECT uid, user_id, title, description, status FROM tasks WHERE user_id = $1 AND deleted = false`
	rows, err := r.db.Query(context.Background(), query, userID)
	if err != nil {
		log.Printf("Failed to get tasks: %v", err)
		return nil, fmt.Errorf("failed to query tasks: %w", err)
	}
	defer rows.Close()

	var tasks []task.Task
	for rows.Next() {
		var tid, uidUserID uuid.UUID
		var title, description string
		var status int

		err := rows.Scan(&tid, &uidUserID, &title, &description, &status)
		if err != nil {
			log.Printf("Failed to scan task row: %v", err)
			return nil, fmt.Errorf("failed to scan task row: %w", err)
		}

		t := task.NewTask(title, description, task.Status(status))
		t.SetID(tid)
		t.SetUserID(uidUserID)
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Rows iteration error: %v", err)
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	log.Printf("Retrieved %d tasks for user %v", len(tasks), userID)
	return tasks, nil
}

func (r *TaskRepository) GetById(userID, id uuid.UUID) (*task.Task, error) {
	query := `SELECT uid, user_id, title, description, status FROM tasks WHERE uid = $1 AND user_id = $2 AND deleted = false`
	row := r.db.QueryRow(context.Background(), query, id, userID)

	var tid, uidUserID uuid.UUID
	var title, description string
	var status int
	err := row.Scan(&tid, &uidUserID, &title, &description, &status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Printf("Task not found by ID: %v", id)
			return nil, fmt.Errorf("task not found: %w", err)
		}
		log.Printf("Error getting task by ID: %v", err)
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	t := task.NewTask(title, description, task.Status(status))
	t.SetID(tid)
	t.SetUserID(uidUserID)

	return &t, nil
}

func (r *TaskRepository) Update(id uuid.UUID, updated task.Task) (*task.Task, error) {
	query := `UPDATE tasks SET title = $1, description = $2, status = $3 WHERE uid = $4 AND deleted = false`
	_, err := r.db.Exec(context.Background(), query,
		updated.Title(), updated.Description(), int(updated.Status()), id)
	if err != nil {
		log.Printf("Failed to update task %v: %v", id, err)
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	log.Printf("Task %v updated successfully", id)
	updated.SetID(id)
	return &updated, nil
}
