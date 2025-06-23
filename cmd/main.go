package main

import (
	"ToDo/internal/config"
	"ToDo/internal/repository/dbstorage"
	inmemory "ToDo/internal/repository/in-memory"
	"ToDo/internal/server"
	"ToDo/internal/server/handlers"
	"ToDo/internal/service"
	"context"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func main() {
	ctx := context.Background()
	log.Println("Starting application...")

	cfg := config.Load()

	dbURL := cfg.DBURL()
	log.Printf("Connecting to DB with URL: %s", dbURL)

	if err := dbstorage.ApplyMigrations(dbURL, "file://migrations"); err != nil {
		log.Printf("Migration failed: %v", err)
	}

	db, err := dbstorage.New(ctx, dbURL)

	var taskRepo service.TaskRepositoryInterface
	var userRepo service.UserRepositoryInterface

	if err != nil {
		log.Printf("DB connection failed: %v, using in-memory storage", err)
		taskRepo = inmemory.NewTaskRepository()
		userRepo = inmemory.NewUserRepository()
	} else {
		defer db.Close(ctx)
		taskRepo = db.TaskRepository()
		userRepo = db.UserRepository()
	}

	taskService := service.NewTaskService(taskRepo)
	userService := service.NewUserService(userRepo)

	taskHandler := handlers.NewTaskHandler(taskService)
	userHandler := handlers.NewUserHandler(userService)

	router := server.NewRouter(taskHandler, userHandler)

	if err := router.Run(cfg.HTTPAddr()); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
