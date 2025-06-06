package dbstorage

import (
	"ToDo/internal/domain/task"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type TaskRepository struct {
	db *pgx.Conn
}

func NewTaskRepository(db *pgx.Conn) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Create(t task.Task) (task.Task, error) {
	query := `INSERT INTO tasks (uid, user_id, title, description, status) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(context.Background(), query,
		t.ID(), t.UserID(), t.Title(), t.Description(), int(t.Status()))
	if err != nil {
		return task.Task{}, fmt.Errorf("failed to insert task: %w", err)
	}
	return t, nil
}

func (r *TaskRepository) GetByTitle(userID uuid.UUID, title string) (*task.Task, error) {
	query := `SELECT uid, user_id, title, description, status FROM tasks WHERE user_id = $1 AND title = $2 LIMIT 1`
	row := r.db.QueryRow(context.Background(), query, userID, title)

	var tid, uidUserID uuid.UUID
	var titleDB, description string
	var status int
	err := row.Scan(&tid, &uidUserID, &titleDB, &description, &status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("task not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get task by title: %w", err)
	}

	t := task.NewTask(titleDB, description, task.Status(status))
	t.SetID(tid)
	t.SetUserID(uidUserID)

	return &t, nil
}

func (r *TaskRepository) GetAll(userID uuid.UUID) ([]task.Task, error) {
	query := `SELECT uid, user_id, title, description, status FROM tasks WHERE user_id = $1`
	rows, err := r.db.Query(context.Background(), query, userID)
	if err != nil {
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
			return nil, fmt.Errorf("failed to scan task row: %w", err)
		}

		t := task.NewTask(title, description, task.Status(status))
		t.SetID(tid)
		t.SetUserID(uidUserID)
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return tasks, nil
}

func (r *TaskRepository) GetById(userID, id uuid.UUID) (*task.Task, error) {
	query := `SELECT uid, user_id, title, description, status FROM tasks WHERE uid = $1 AND user_id = $2`
	row := r.db.QueryRow(context.Background(), query, id, userID)

	var tid, uidUserID uuid.UUID
	var title, description string
	var status int
	err := row.Scan(&tid, &uidUserID, &title, &description, &status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("task not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get task: %w", err)
	}

	t := task.NewTask(title, description, task.Status(status))
	t.SetID(tid)
	t.SetUserID(uidUserID)

	return &t, nil
}

func (r *TaskRepository) Update(id uuid.UUID, updated task.Task) (*task.Task, error) {
	query := `UPDATE tasks SET title = $1, description = $2, status = $3 WHERE uid = $4`
	_, err := r.db.Exec(context.Background(), query,
		updated.Title(), updated.Description(), int(updated.Status()), id)
	if err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	updated.SetID(id)
	return &updated, nil
}

func (r *TaskRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM tasks WHERE uid = $1`
	_, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}
