package task

import "github.com/google/uuid"

// CreateTaskDTO используется для создания новой задачи из входных данных JSON.
type CreateTaskDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      int    `json:"status"`
}

// DTO представляет собой транспортный объект задачи для ответа клиенту.
type DTO struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	UserID      uuid.UUID `json:"user_id"`
}
