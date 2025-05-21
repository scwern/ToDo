package task

type CreateTaskDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      int    `json:"status"`
}

type DTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}
