package task

import (
	"ToDo/internal/domain/task"
)

func ToTask(dto CreateTaskDTO) task.Task {
	status := task.Status(dto.Status)
	if status == "" {
		status = task.StatusNew
	}

	return task.NewTask(dto.Title, dto.Description, status)
}

func ToTaskDTO(t task.Task) DTO {
	return DTO{
		ID:          t.ID().String(),
		Title:       t.Title(),
		Description: t.Description(),
		Status:      string(t.Status()),
	}
}
