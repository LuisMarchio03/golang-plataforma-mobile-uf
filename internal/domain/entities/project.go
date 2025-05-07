package entities

type ProjectStatus string

const (
	StatusOpen       ProjectStatus = "open"
	StatusInProgress ProjectStatus = "in_progress"
	StatusCompleted  ProjectStatus = "completed"
)

type Project struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Status      ProjectStatus `json:"status"`
	StartDate   string        `json:"start_date"`
	EndDate     string        `json:"end_date"`
	CreatedBy   string        `json:"created_by"`
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
}
