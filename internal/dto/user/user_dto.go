package user

// CreateUserDTO используется для создания нового пользователя из входных данных JSON.
type CreateUserDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// DTO представляет собой транспортный объект пользователя для ответа клиенту.
type DTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
