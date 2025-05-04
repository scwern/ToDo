package main

import (
	"ToDo/internal/repository"
	"ToDo/internal/server"
	"ToDo/internal/server/handlers"
	"ToDo/internal/service"
)

func main() {
	taskRepo := repository.NewTaskRepository()
	userRepo := repository.NewUserRepository()

	taskService := service.NewTaskService(taskRepo)
	userService := service.NewUserService(userRepo)

	taskHandler := handlers.NewTaskHandler(taskService)
	userHandler := handlers.NewUserHandler(userService)

	r := server.NewRouter(taskHandler, userHandler)

	r.Run(":8080")
}
