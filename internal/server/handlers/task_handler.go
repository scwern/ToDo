package handlers

import (
	taskdto "ToDo/internal/dto/task"
	"ToDo/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// TaskHandler предоставляет HTTP-обработчики для операций с задачами.
// @Description Обрабатывает CRUD операции для задач пользователей
type TaskHandler struct {
	service *service.TaskService
}

// NewTaskHandler создает новый экземпляр TaskHandler.
// @Summary Создать обработчик задач
func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

// GetAll обрабатывает HTTP-запрос на получение всех задач пользователя.
// @Summary Получить все задачи
// @Description Возвращает список задач для текущего пользователя
// @Tags Tasks
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} taskdto.DTO
// @Failure 401 {object} object "Не авторизован"
// @Failure 500 {object} object "Внутренняя ошибка сервера"
// @Router /tasks [get]
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
// @Summary Получить задачу по ID
// @Description Возвращает задачу по указанному идентификатору
// @Tags Tasks
// @Produce json
// @Param id path string true "UUID задачи"
// @Security ApiKeyAuth
// @Success 200 {object} taskdto.DTO
// @Failure 400 {object} object "Неверный формат UUID"
// @Failure 404 {object} object "Задача не найдена"
// @Router /tasks/{id} [get]
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
// @Summary Создать новую задачу
// @Description Создает новую задачу для текущего пользователя
// @Tags Tasks
// @Accept json
// @Produce json
// @Param input body taskdto.CreateTaskDTO true "Данные задачи"
// @Security ApiKeyAuth
// @Success 201 {object} taskdto.DTO
// @Failure 400 {object} object "Неверные входные данные"
// @Failure 409 {object} object "Задача с таким названием уже существует"
// @Router /tasks [post]
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
// @Summary Обновить задачу
// @Description Обновляет существующую задачу
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path string true "UUID задачи"
// @Param input body taskdto.CreateTaskDTO true "Обновленные данные задачи"
// @Security ApiKeyAuth
// @Success 200 {object} taskdto.DTO
// @Failure 400 {object} object "Неверные входные данные"
// @Failure 404 {object} object "Задача не найдена"
// @Router /tasks/{id} [put]
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
// @Summary Удалить задачу
// @Description Удаляет задачу по указанному идентификатору
// @Tags Tasks
// @Param id path string true "UUID задачи"
// @Security ApiKeyAuth
// @Success 204 "Задача успешно удалена"
// @Failure 400 {object} object "Неверный формат UUID"
// @Failure 404 {object} object "Задача не найдена"
// @Router /tasks/{id} [delete]
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
