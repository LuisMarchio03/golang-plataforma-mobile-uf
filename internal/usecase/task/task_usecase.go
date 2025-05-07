package usecase

import (
	"context"

	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/repository"
)

type TaskUseCase interface {
	Create(ctx context.Context, input CreateTaskInput) (*repository.Task, error)
	GetByID(ctx context.Context, id string) (*repository.Task, error)
	List(ctx context.Context) ([]repository.Task, error)
	UpdateStatus(ctx context.Context, input UpdateTaskStatusInput) error
	AssignTask(ctx context.Context, input AssignTaskInput) error
	ListByProject(ctx context.Context, projectID string) ([]repository.Task, error)
	ListByAssignee(ctx context.Context, userID string) ([]repository.Task, error)
}

type CreateTaskInput struct {
	ProjectID   string `json:"project_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateTaskStatusInput struct {
	ID     string                `json:"id" validate:"required"`
	Status repository.TaskStatus `json:"status" validate:"required,oneof=pending in_progress completed"`
}

type AssignTaskInput struct {
	TaskID     string `json:"task_id" validate:"required"`
	AssignedTo string `json:"assigned_to" validate:"required"`
}

func NewTaskUseCaseImpl(taskRepo repository.TaskRepository, projectRepo repository.ProjectRepository) TaskUseCase {
	return &taskUseCase{
		taskRepo:    taskRepo,
		projectRepo: projectRepo,
	}
}
