package main

import (
	"ToDo/internal/repository"
	"ToDo/internal/server"
	"ToDo/internal/server/handlers"
	"ToDo/internal/service"
)

func main() {
	taskRepo := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)

	r := server.NewRouter(taskHandler)
	r.Run(":8080")
}
