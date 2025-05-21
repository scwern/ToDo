package db_storage

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
	query := `INSERT INTO tasks (uid, title, description, status) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(context.Background(), query,
		t.ID(), t.Title(), t.Description(), int(t.Status()))
	if err != nil {
		return task.Task{}, fmt.Errorf("failed to insert task: %w", err)
	}
	return t, nil
}

func (r *TaskRepository) GetByTitle(title string) (*task.Task, error) {
	query := `SELECT uid, title, description, status FROM tasks WHERE title = $1 LIMIT 1`
	row := r.db.QueryRow(context.Background(), query, title)

	var tid uuid.UUID
	var titleDB, description string
	var status int
	err := row.Scan(&tid, &titleDB, &description, &status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("task not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get task by title: %w", err)
	}

	t := task.NewTask(titleDB, description, task.Status(status))
	t.SetID(tid)

	return &t, nil
}

func (r *TaskRepository) GetAll() ([]task.Task, error) {
	query := `SELECT uid, title, description, status FROM tasks`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to query tasks: %w", err)
	}
	defer rows.Close()

	var tasks []task.Task
	for rows.Next() {
		var tid uuid.UUID
		var title, description string
		var status int

		err := rows.Scan(&tid, &title, &description, &status)
		if err != nil {
			return nil, fmt.Errorf("failed to scan task row: %w", err)
		}

		t := task.NewTask(title, description, task.Status(status))
		t.SetID(tid)
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return tasks, nil
}

func (r *TaskRepository) GetById(id uuid.UUID) (*task.Task, error) {
	var t task.Task
	var status int

	query := `SELECT uid, title, description, status FROM tasks WHERE uid = $1`
	row := r.db.QueryRow(context.Background(), query, id)

	var tid uuid.UUID
	var title, description string
	err := row.Scan(&tid, &title, &description, &status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &t, fmt.Errorf("task not found: %w", err)
		}
		return &t, fmt.Errorf("failed to get task: %w", err)
	}

	t = task.NewTask(title, description, task.Status(status))
	t.SetID(tid)

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
