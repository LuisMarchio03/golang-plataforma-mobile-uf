package entities

type ApplicationStatus string

const (
	ApplicationStatusPending  ApplicationStatus = "pending"
	ApplicationStatusApproved ApplicationStatus = "approved"
	ApplicationStatusRejected ApplicationStatus = "rejected"
)

type ProjectApplication struct {
	ID        string            `json:"id"`
	ProjectID string            `json:"project_id"`
	UserID    string            `json:"user_id"`
	Status    ApplicationStatus `json:"status"`
	CreatedAt string            `json:"created_at"`
	UpdatedAt string            `json:"updated_at"`
}
