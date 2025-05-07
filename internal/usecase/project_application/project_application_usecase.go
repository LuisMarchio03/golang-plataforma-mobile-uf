package usecase

import (
	"context"

	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/repository"
)

type ProjectApplicationUseCase interface {
	Create(ctx context.Context, input CreateProjectApplicationInput) (*repository.ProjectApplication, error)
	GetByID(ctx context.Context, id string) (*repository.ProjectApplication, error)
	List(ctx context.Context) ([]repository.ProjectApplication, error)
	UpdateStatus(ctx context.Context, input UpdateApplicationStatusInput) error
	ListByProject(ctx context.Context, projectID string) ([]repository.ProjectApplication, error)
	ListByUser(ctx context.Context, userID string) ([]repository.ProjectApplication, error)
}

type CreateProjectApplicationInput struct {
	ProjectID string `json:"project_id" validate:"required"`
	UserID    string `json:"user_id" validate:"required"`
	Message   string `json:"message" validate:"required"`
}

type UpdateApplicationStatusInput struct {
	ID     string                       `json:"id" validate:"required"`
	Status repository.ApplicationStatus `json:"status" validate:"required,oneof=pending approved rejected"`
}
