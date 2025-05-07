package handlers

import (
	"ToDo/internal/domain/task"
	taskdto "ToDo/internal/dto/task"
	"ToDo/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetAll(c *gin.Context) {
	tasks := h.service.GetAll()

	var taskDTOs []taskdto.DTO
	for _, t := range tasks {
		taskDTOs = append(taskDTOs, taskdto.ToTaskDTO(t))
	}

	c.JSON(http.StatusOK, taskDTOs)
}

func (h *TaskHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	fmt.Println("Received ID:", idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	fmt.Println("Parsed ID:", id)

	task, err := h.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	fmt.Println("Found task:", task)

	c.JSON(http.StatusOK, taskdto.ToTaskDTO(*task))
}

func (h *TaskHandler) Create(c *gin.Context) {
	var input taskdto.CreateTaskDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingTask, _ := h.service.GetByTitle(input.Title)
	if existingTask != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Task with this title already exists"})
		return
	}

	newTask := taskdto.ToTask(input)
	created := h.service.Create(newTask)
	response := taskdto.ToTaskDTO(created)

	c.JSON(http.StatusCreated, response)
}

func (h *TaskHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	var t task.Task
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.service.Update(id, t)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (h *TaskHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
