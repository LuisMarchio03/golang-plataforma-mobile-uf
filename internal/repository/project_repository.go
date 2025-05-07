package repository

import (
	"context"
	"time"
)

type ProjectStatus string

const (
	ProjectStatusOpen       ProjectStatus = "open"
	ProjectStatusInProgress ProjectStatus = "in_progress"
	ProjectStatusCompleted  ProjectStatus = "completed"
)

type Project struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Status      ProjectStatus `json:"status"`
	CreatedBy   string        `json:"created_by"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type ProjectRepository interface {
	Repository[Project]
	FindByStatus(ctx context.Context, status ProjectStatus) ([]Project, error)
	FindByCreator(ctx context.Context, creatorID string) ([]Project, error)
}
