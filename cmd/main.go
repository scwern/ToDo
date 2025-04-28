package main

import (
	"ToDo/internal/handler"
	"ToDo/internal/repository"
	"ToDo/internal/router"
	"ToDo/internal/service"
)

func main() {
	taskRepo := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	r := router.NewRouter(taskHandler)

	r.Run(":8080")
}
