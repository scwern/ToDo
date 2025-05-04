package main

import (
	"ToDo/internal/handler"
	"ToDo/internal/repository"
	"ToDo/internal/router"
	"ToDo/internal/service"
)

func main() {
	taskRepo := repository.NewTaskRepository()
	userRepo := repository.NewUserRepository()

	taskService := service.NewTaskService(taskRepo)
	userService := service.NewUserService(userRepo)

	taskHandler := handler.NewTaskHandler(taskService)
	userHandler := handler.NewUserHandler(userService)

	r := router.NewRouter(taskHandler, userHandler)

	r.Run(":8080")
}
