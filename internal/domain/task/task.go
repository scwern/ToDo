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
	userID      uuid.UUID `json:"userID"`
	deleted     bool      `json:"-"`
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

func (t *Task) UserID() uuid.UUID {
	return t.userID
}

func (t *Task) IsDeleted() bool {
	return t.deleted
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

func (t *Task) SetUserID(id uuid.UUID) {
	t.userID = id
}

func (t *Task) SetDeleted(d bool) {
	t.deleted = d
}
