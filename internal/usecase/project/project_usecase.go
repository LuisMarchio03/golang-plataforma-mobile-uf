package usecase

import (
	"context"

	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/repository"
)

type ProjectUseCase interface {
	Create(ctx context.Context, input CreateProjectInput) (*repository.Project, error)
	GetByID(ctx context.Context, id string) (*repository.Project, error)
	List(ctx context.Context) ([]repository.Project, error)
	Update(ctx context.Context, input UpdateProjectInput) (*repository.Project, error)
	Delete(ctx context.Context, id string) error
	UpdateStatus(ctx context.Context, input UpdateProjectStatusInput) error
	ListByStatus(ctx context.Context, status repository.ProjectStatus) ([]repository.Project, error)
}

type CreateProjectInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	CreatedBy   string `json:"created_by" validate:"required"`
}

type UpdateProjectInput struct {
	ID          string `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateProjectStatusInput struct {
	ID     string                   `json:"id" validate:"required"`
	Status repository.ProjectStatus `json:"status" validate:"required,oneof=open in_progress completed"`
}
