package db_storage

import (
	"ToDo/internal/domain/user"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(u user.User) (user.User, error) {
	query := `INSERT INTO users (uid, name, email, password) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(context.Background(), query,
		u.ID(), u.Name(), u.Email(), u.Password())
	if err != nil {
		return user.User{}, fmt.Errorf("failed to insert user: %w", err)
	}
	return u, nil
}

func (r *UserRepository) GetByEmail(email string) (*user.User, error) {
	query := `SELECT uid, name, email, password FROM users WHERE email = $1`
	row := r.db.QueryRow(context.Background(), query, email)

	var uid uuid.UUID
	var name, emailDB, password string
	err := row.Scan(&uid, &name, &emailDB, &password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	u := user.NewUser(name, emailDB, password)
	u.SetID(uid)

	return &u, nil
}

func (r *UserRepository) GetAll() ([]user.User, error) {
	query := `SELECT uid, name, email, password FROM users`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var uid uuid.UUID
		var name, email, password string

		err := rows.Scan(&uid, &name, &email, &password)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user row: %w", err)
		}

		u := user.NewUser(name, email, password)
		u.SetID(uid)

		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return users, nil
}

func (r *UserRepository) GetById(id uuid.UUID) (*user.User, error) {
	var u user.User

	query := `SELECT uid, name, email, password FROM users WHERE id = $1`
	row := r.db.QueryRow(context.Background(), query, id)

	var uid uuid.UUID
	var name, email, password string
	err := row.Scan(&uid, &name, &email, &password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &u, fmt.Errorf("user not found: %w", err)
		}
		return &u, fmt.Errorf("failed to get user: %w", err)
	}

	u = user.NewUser(name, email, password)
	u.SetID(uid)

	return &u, nil
}

func (r *UserRepository) Update(id uuid.UUID, updated user.User) (*user.User, error) {
	query := `UPDATE users SET name = $1, email = $2, password = $3 WHERE uid = $4`
	_, err := r.db.Exec(context.Background(), query,
		updated.Name(), updated.Email(), updated.Password(), id)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	updated.SetID(id)
	return &updated, nil
}

func (r *UserRepository) Delete(id uuid.UUID) error {
	query := `DELETE FROM users WHERE uid = $1`
	_, err := r.db.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
