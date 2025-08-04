package dbstorage

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgx/v5"
	"log"
)

type DBStorage struct {
	db *pgx.Conn
}

func New(ctx context.Context, addr string) (*DBStorage, error) {
	log.Printf("Connecting to DB with URL: %s", addr)
	conn, err := pgx.Connect(ctx, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	fmt.Println("Connected to database")
	return &DBStorage{db: conn}, nil
}

func (db *DBStorage) Close(ctx context.Context) error {
	if db.TaskRepository() != nil {
		db.TaskRepository().Close()
	}
	return db.db.Close(ctx)
}

func (db *DBStorage) TaskRepository() *TaskRepository {
	return NewTaskRepository(db.db)
}

func (db *DBStorage) UserRepository() *UserRepository {
	return NewUserRepository(db.db)
}

func ApplyMigrations(addr, migrationPath string) error {
	m, err := migrate.New(migrationPath, addr)
	if err != nil {
		return fmt.Errorf("failed to initialize migrate: %w", err)
	}
	defer func() { err, _ = m.Close() }()

	err = m.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("No changes to apply")
			return nil
		}
		return fmt.Errorf("failed to apply migrations: %w", err)
	}

	fmt.Println("Migrations applied successfully")
	return nil
}
