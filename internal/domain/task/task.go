package task

import "github.com/google/uuid"

type Status int

const (
	StatusNew Status = iota
	StatusInProgress
	StatusDone
)

func (s Status) String() string {
	switch s {
	case StatusNew:
		return "Новая"
	case StatusInProgress:
		return "В процессе"
	case StatusDone:
		return "Завершена"
	default:
		return "Неизвестно"
	}
}

type Task struct {
	id          uuid.UUID `json:"id"`
	title       string    `json:"title"`
	description string    `json:"description"`
	status      Status    `json:"status"`
}

func NewTask(title, description string, status Status) Task {
	if status == 0 {
		status = StatusNew
	}
	return Task{
		id:          uuid.New(),
		title:       title,
		description: description,
		status:      status,
	}
}

func (t *Task) ID() uuid.UUID {
	return t.id
}

func (t *Task) Title() string {
	return t.title
}

func (t *Task) Description() string {
	return t.description
}

func (t *Task) Status() Status {
	return t.status
}

func (t *Task) SetTitle(title string) {
	t.title = title
}

func (t *Task) SetDescription(desc string) {
	t.description = desc
}

func (t *Task) SetStatus(status Status) {
	t.status = status
}

func (t *Task) SetID(id uuid.UUID) {
	t.id = id
}
