package repository

import (
	"context"
	"time"
)

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
	Message   string            `json:"message"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type ProjectApplicationRepository interface {
	Repository[ProjectApplication]
	FindByProject(ctx context.Context, projectID string) ([]ProjectApplication, error)
	FindByUser(ctx context.Context, userID string) ([]ProjectApplication, error)
	FindByStatus(ctx context.Context, status ApplicationStatus) ([]ProjectApplication, error)
}
