package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/repository"
	"github.com/google/uuid"
)

type projectApplicationUseCase struct {
	applicationRepo repository.ProjectApplicationRepository
	projectRepo     repository.ProjectRepository
	userRepo        repository.UserRepository
}

func NewProjectApplicationUseCase(
	applicationRepo repository.ProjectApplicationRepository,
	projectRepo repository.ProjectRepository,
	userRepo repository.UserRepository,
) ProjectApplicationUseCase {
	return &projectApplicationUseCase{
		applicationRepo: applicationRepo,
		projectRepo:     projectRepo,
		userRepo:        userRepo,
	}
}

func (uc *projectApplicationUseCase) Create(ctx context.Context, input CreateProjectApplicationInput) (*repository.ProjectApplication, error) {
	// Verifica se o projeto existe
	project, err := uc.projectRepo.FindByID(ctx, input.ProjectID)
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar projeto: %v", err)
	}
	if project == nil {
		return nil, fmt.Errorf("projeto não encontrado")
	}

	// Verifica se o usuário existe
	user, err := uc.userRepo.FindByID(ctx, input.UserID)
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar usuário: %v", err)
	}
	if user == nil {
		return nil, fmt.Errorf("usuário não encontrado")
	}

	// Verifica se já existe uma candidatura pendente deste usuário para este projeto
	existingApplications, err := uc.applicationRepo.FindByProject(ctx, input.ProjectID)
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar candidaturas existentes: %v", err)
	}

	for _, app := range existingApplications {
		if app.UserID == input.UserID && app.Status == repository.ApplicationStatusPending {
			return nil, fmt.Errorf("já existe uma candidatura pendente para este projeto")
		}
	}

	application := &repository.ProjectApplication{
		ID:        uuid.New().String(),
		ProjectID: input.ProjectID,
		UserID:    input.UserID,
		Message:   input.Message,
		Status:    repository.ApplicationStatusPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.applicationRepo.Create(ctx, *application); err != nil {
		return nil, fmt.Errorf("erro ao criar candidatura: %v", err)
	}

	return application, nil
}

func (uc *projectApplicationUseCase) GetByID(ctx context.Context, id string) (*repository.ProjectApplication, error) {
	application, err := uc.applicationRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar candidatura: %v", err)
	}
	if application == nil {
		return nil, fmt.Errorf("candidatura não encontrada")
	}
	return application, nil
}

func (uc *projectApplicationUseCase) List(ctx context.Context) ([]repository.ProjectApplication, error) {
	applications, err := uc.applicationRepo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar candidaturas: %v", err)
	}
	return applications, nil
}

func (uc *projectApplicationUseCase) UpdateStatus(ctx context.Context, input UpdateApplicationStatusInput) error {
	application, err := uc.applicationRepo.FindByID(ctx, input.ID)
	if err != nil {
		return fmt.Errorf("erro ao buscar candidatura: %v", err)
	}
	if application == nil {
		return fmt.Errorf("candidatura não encontrada")
	}

	// Verifica se a candidatura já foi processada
	if application.Status != repository.ApplicationStatusPending {
		return fmt.Errorf("candidatura já foi processada anteriormente")
	}

	application.Status = input.Status
	application.UpdatedAt = time.Now()

	if err := uc.applicationRepo.Update(ctx, *application); err != nil {
		return fmt.Errorf("erro ao atualizar status da candidatura: %v", err)
	}

	return nil
}

func (uc *projectApplicationUseCase) ListByProject(ctx context.Context, projectID string) ([]repository.ProjectApplication, error) {
	applications, err := uc.applicationRepo.FindByProject(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar candidaturas do projeto: %v", err)
	}
	return applications, nil
}

func (uc *projectApplicationUseCase) ListByUser(ctx context.Context, userID string) ([]repository.ProjectApplication, error) {
	applications, err := uc.applicationRepo.FindByUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar candidaturas do usuário: %v", err)
	}
	return applications, nil
}
