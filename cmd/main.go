package main

import (
	"ToDo/internal/config"
	"ToDo/internal/repository/dbstorage"
	inmemory "ToDo/internal/repository/in-memory"
	"ToDo/internal/server"
	"ToDo/internal/server/handlers"
	"ToDo/internal/service"
	"context"
	"errors"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
		defer func() {
			if err := db.Close(ctx); err != nil {
				log.Printf("Failed to close DB: %v", err)
			}
		}()
		taskRepo = db.TaskRepository()
		userRepo = db.UserRepository()
	}

	taskService := service.NewTaskService(taskRepo)
	userService := service.NewUserService(userRepo)

	taskHandler := handlers.NewTaskHandler(taskService)
	userHandler := handlers.NewUserHandler(userService)

	router := server.NewRouter(taskHandler, userHandler)

	srv := &http.Server{
		Addr:    cfg.HTTPAddr(),
		Handler: router,
	}

	go func() {
		log.Printf("Server is running on %s", cfg.HTTPAddr())
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server Shutdown Failed: %v", err)
	}

	log.Println("Server exited properly")
}
