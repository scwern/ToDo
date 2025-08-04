package handlers

import (
	taskdto "ToDo/internal/dto/task"
	"ToDo/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// TaskHandler предоставляет HTTP-обработчики для операций с задачами.
type TaskHandler struct {
	service *service.TaskService
}

// NewTaskHandler создает новый экземпляр TaskHandler.
func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// GetAll обрабатывает HTTP-запрос на получение всех задач пользователя.
func (h *TaskHandler) GetAll(c *gin.Context) {
	userIDStr, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing user_id cookie"})
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id UUID"})
		return
	}

	tasks, err := h.service.GetAll(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var taskDTOs []taskdto.DTO
	for _, t := range tasks {
		taskDTOs = append(taskDTOs, taskdto.ToTaskDTO(t))
	}

	c.JSON(http.StatusOK, taskDTOs)
}

// GetByID обрабатывает HTTP-запрос на получение задачи по UUID.
func (h *TaskHandler) GetByID(c *gin.Context) {
	userIDStr, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing user_id cookie"})
		return
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id UUID"})
		return
	}

	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	task, err := h.service.GetById(userID, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, taskdto.ToTaskDTO(*task))
}

// Create обрабатывает HTTP-запрос на создание новой задачи.
func (h *TaskHandler) Create(c *gin.Context) {
	userIDStr, err := c.Cookie("user_id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing user_id cookie"})
		return
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id UUID"})
		return
	}

	var input taskdto.CreateTaskDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingTask, _ := h.service.GetByTitle(userID, input.Title)
	if existingTask != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Task with this title already exists"})
		return
	}

	newTask := taskdto.ToTask(input)
	newTask.SetUserID(userID)

	created, err := h.service.Create(newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := taskdto.ToTaskDTO(created)

	c.JSON(http.StatusCreated, response)
}

// Update обрабатывает HTTP-запрос на обновление задачи по UUID.
func (h *TaskHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	var input taskdto.CreateTaskDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask := taskdto.ToTask(input)
	updated, err := h.service.Update(id, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, taskdto.ToTaskDTO(*updated))
}

// Delete обрабатывает HTTP-запрос на удаление задачи по UUID.
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
