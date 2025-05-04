package router

import (
	"ToDo/internal/handler"
	"ToDo/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(taskHandler *handler.TaskHandler, userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/login", userHandler.Login)
	r.POST("/users", userHandler.Create)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())

	auth.GET("/profile", userHandler.Profile)

	tasks := auth.Group("/tasks")
	{
		tasks.GET("", taskHandler.GetAll)
		tasks.GET("/:id", taskHandler.GetByID)
		tasks.POST("", taskHandler.Create)
		tasks.PUT("/:id", taskHandler.Update)
		tasks.DELETE("/:id", taskHandler.Delete)
	}

	users := auth.Group("/users")
	{
		users.GET("", userHandler.GetAll)
		users.GET("/:id", userHandler.GetByID)
		users.PUT("/:id", userHandler.Update)
		users.DELETE("/:id", userHandler.Delete)
	}

	return r
}
