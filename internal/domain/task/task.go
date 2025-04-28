package task

type Status string

const (
	StatusNew        Status = "New"
	StatusInProgress Status = "InProgress"
	StatusDone       Status = "Done"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}
