package server

import (
	"ToDo/internal/server/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewRouter(taskHandler *handlers.TaskHandler, userHandler *handlers.UserHandler) *gin.Engine {
	r := gin.Default()

	tasks := r.Group("/tasks")
	{
		tasks.GET("", func(c *gin.Context) {
			fmt.Println("Request to GET /tasks")
			taskHandler.GetAll(c)
		})
		tasks.GET("/:id", func(c *gin.Context) {
			fmt.Println("Request to GET /tasks/:id")
			taskHandler.GetByID(c)
		})
		tasks.POST("", func(c *gin.Context) {
			fmt.Println("Request to POST /tasks")
			taskHandler.Create(c)
		})
		tasks.PUT("/:id", func(c *gin.Context) {
			fmt.Println("Request to PUT /tasks/:id")
			taskHandler.Update(c)
		})
		tasks.DELETE("/:id", func(c *gin.Context) {
			fmt.Println("Request to DELETE /tasks/:id")
			taskHandler.Delete(c)
		})
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
