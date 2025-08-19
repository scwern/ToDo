package task

import (
	"ToDo/internal/domain/task"
)

// ToTask преобразует CreateTaskDTO в доменную сущность task.Task.
func ToTask(dto CreateTaskDTO) task.Task {
	status := task.Status(dto.Status)
	if status == 0 {
		status = task.StatusNew
	}

	return task.NewTask(dto.Title, dto.Description, status)
}

// ToTaskDTO преобразует доменную сущность task.Task в DTO для передачи клиенту.
func ToTaskDTO(t task.Task) DTO {
	return DTO{
		ID:          t.ID().String(),
		Title:       t.Title(),
		Description: t.Description(),
		Status:      int(t.Status()),
		UserID:      t.UserID(),
	}
}
