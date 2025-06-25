package server

import (
	"ToDo/internal/middleware"
	"ToDo/internal/server/handlers"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func NewRouter(taskHandler *handlers.TaskHandler, userHandler *handlers.UserHandler) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())

	r.Use(gzip.Gzip(gzip.BestSpeed))
	r.Use(middleware.GzipRequestMiddleware())

	tasks := r.Group("/tasks")
	{
		tasks.GET("", taskHandler.GetAll)
		tasks.GET("/:id", taskHandler.GetByID)
		tasks.POST("", taskHandler.Create)
		tasks.PUT("/:id", taskHandler.Update)
		tasks.DELETE("/:id", taskHandler.Delete)
	}

	users := r.Group("/users")
	{
		users.GET("", userHandler.GetAll)
		users.GET("/:id", userHandler.GetByID)
		users.POST("", userHandler.Create)
		users.PUT("/:id", userHandler.Update)
		users.DELETE("/:id", userHandler.Delete)
	}

	return r
}
