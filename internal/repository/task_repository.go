package repository

import (
	"context"
	"time"
)

type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
)

type Task struct {
	ID          string     `json:"id"`
	ProjectID   string     `json:"project_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	AssignedTo  string     `json:"assigned_to"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type TaskRepository interface {
	Repository[Task]
	FindByProject(ctx context.Context, projectID string) ([]Task, error)
	FindByAssignee(ctx context.Context, assigneeID string) ([]Task, error)
	FindByStatus(ctx context.Context, status TaskStatus) ([]Task, error)
}
