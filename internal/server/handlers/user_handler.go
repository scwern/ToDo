package handlers

import (
	"ToDo/internal/domain/user"
	userdto "ToDo/internal/dto/user"
	"ToDo/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// UserHandler предоставляет HTTP-обработчики для операций с пользователями.
// @Description Обрабатывает CRUD операции для пользователей системы
type UserHandler struct {
	service *service.UserService
}

// NewUserHandler создает новый экземпляр UserHandler.
// @Summary Создать обработчик пользователей
func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// GetAll обрабатывает HTTP-запрос на получение всех пользователей.
// @Summary Получить всех пользователей
// @Description Возвращает список всех зарегистрированных пользователей
// @Tags Users
// @Produce json
// @Success 200 {array} userdto.DTO
// @Failure 500 {object} object "Внутренняя ошибка сервера"
// @Router /users [get]
func (h *UserHandler) GetAll(c *gin.Context) {
	users, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userDTOs []userdto.DTO
	for _, u := range users {
		userDTOs = append(userDTOs, userdto.ToUserDTO(u))
	}

	c.JSON(http.StatusOK, userDTOs)
}

// GetByID обрабатывает HTTP-запрос на получение пользователя по UUID.
// @Summary Получить пользователя по ID
// @Description Возвращает пользователя по указанному идентификатору
// @Tags Users
// @Produce json
// @Param id path string true "UUID пользователя"
// @Success 200 {object} userdto.DTO
// @Failure 400 {object} object "Неверный формат UUID"
// @Failure 404 {object} object "Пользователь не найден"
// @Router /users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	fmt.Println("Received ID:", idStr)

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	fmt.Println("Parsed ID:", id)

	user, err := h.service.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, userdto.ToUserDTO(*user))
}

// Create обрабатывает HTTP-запрос на создание нового пользователя.
// @Summary Создать пользователя
// @Description Регистрирует нового пользователя в системе
// @Tags Users
// @Accept json
// @Produce json
// @Param input body userdto.CreateUserDTO true "Данные пользователя"
// @Success 201 {object} userdto.DTO
// @Failure 400 {object} object "Неверные входные данные"
// @Failure 409 {object} object "Пользователь с таким email уже существует"
// @Router /users [post]
func (h *UserHandler) Create(c *gin.Context) {
	var input userdto.CreateUserDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.service.GetByEmail(input.Email)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User with this email already exists"})
		return
	}

	newUser := userdto.ToUser(input)
	created, err := h.service.Create(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("user_id", created.ID().String(), 3600*24, "/", "localhost", false, true)
	response := userdto.ToUserDTO(created)
	c.JSON(http.StatusCreated, response)
}

// Update обрабатывает HTTP-запрос на обновление пользователя по UUID.
// @Summary Обновить пользователя
// @Description Обновляет данные пользователя по указанному идентификатору
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "UUID пользователя"
// @Param input body user.User true "Обновленные данные пользователя"
// @Success 200 {object} user.User
// @Failure 400 {object} object "Неверные входные данные"
// @Failure 404 {object} object "Пользователь не найден"
// @Router /users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	var u user.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.service.Update(id, u)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, updated)
}

// Delete обрабатывает HTTP-запрос на удаление пользователя по UUID.
// @Summary Удалить пользователя
// @Description Удаляет пользователя по указанному идентификатору
// @Tags Users
// @Param id path string true "UUID пользователя"
// @Success 204 "Пользователь успешно удален"
// @Failure 400 {object} object "Неверный формат UUID"
// @Failure 404 {object} object "Пользователь не найден"
// @Router /users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
