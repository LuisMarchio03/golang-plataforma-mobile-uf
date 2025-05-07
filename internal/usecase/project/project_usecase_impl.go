package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/repository"
	"github.com/google/uuid"
)

type projectUseCase struct {
	projectRepo repository.ProjectRepository
}

func NewProjectUseCase(projectRepo repository.ProjectRepository) ProjectUseCase {
	return &projectUseCase{
		projectRepo: projectRepo,
	}
}

func (uc *projectUseCase) Create(ctx context.Context, input CreateProjectInput) (*repository.Project, error) {
	project := &repository.Project{
		ID:          uuid.New().String(),
		Title:       input.Title,
		Description: input.Description,
		Status:      repository.ProjectStatusOpen,
		CreatedBy:   input.CreatedBy,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := uc.projectRepo.Create(ctx, *project); err != nil {
		return nil, fmt.Errorf("erro ao criar projeto: %v", err)
	}

	return project, nil
}

func (uc *projectUseCase) GetByID(ctx context.Context, id string) (*repository.Project, error) {
	project, err := uc.projectRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar projeto: %v", err)
	}
	if project == nil {
		return nil, fmt.Errorf("projeto não encontrado")
	}
	return project, nil
}

func (uc *projectUseCase) List(ctx context.Context) ([]repository.Project, error) {
	projects, err := uc.projectRepo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar projetos: %v", err)
	}
	return projects, nil
}

func (uc *projectUseCase) Update(ctx context.Context, input UpdateProjectInput) (*repository.Project, error) {
	existingProject, err := uc.projectRepo.FindByID(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar projeto: %v", err)
	}
	if existingProject == nil {
		return nil, fmt.Errorf("projeto não encontrado")
	}

	project := &repository.Project{
		ID:          input.ID,
		Title:       input.Title,
		Description: input.Description,
		Status:      existingProject.Status,
		CreatedBy:   existingProject.CreatedBy,
		CreatedAt:   existingProject.CreatedAt,
		UpdatedAt:   time.Now(),
	}

	if err := uc.projectRepo.Update(ctx, *project); err != nil {
		return nil, fmt.Errorf("erro ao atualizar projeto: %v", err)
	}

	return project, nil
}

func (uc *projectUseCase) Delete(ctx context.Context, id string) error {
	existingProject, err := uc.projectRepo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("erro ao buscar projeto: %v", err)
	}
	if existingProject == nil {
		return fmt.Errorf("projeto não encontrado")
	}

	if err := uc.projectRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("erro ao deletar projeto: %v", err)
	}

	return nil
}

func (uc *projectUseCase) UpdateStatus(ctx context.Context, input UpdateProjectStatusInput) error {
	existingProject, err := uc.projectRepo.FindByID(ctx, input.ID)
	if err != nil {
		return fmt.Errorf("erro ao buscar projeto: %v", err)
	}
	if existingProject == nil {
		return fmt.Errorf("projeto não encontrado")
	}

	existingProject.Status = input.Status
	existingProject.UpdatedAt = time.Now()

	if err := uc.projectRepo.Update(ctx, *existingProject); err != nil {
		return fmt.Errorf("erro ao atualizar status do projeto: %v", err)
	}

	return nil
}

func (uc *projectUseCase) ListByStatus(ctx context.Context, status repository.ProjectStatus) ([]repository.Project, error) {
	projects, err := uc.projectRepo.FindByStatus(ctx, status)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar projetos por status: %v", err)
	}
	return projects, nil
}
