package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/repository"
	"github.com/google/uuid"
)

type taskUseCase struct {
	taskRepo    repository.TaskRepository
	projectRepo repository.ProjectRepository
}

// Delete implements TaskUseCase.
func (uc *taskUseCase) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

func NewTaskUseCase(taskRepo repository.TaskRepository, projectRepo repository.ProjectRepository) TaskUseCase {
	return &taskUseCase{
		taskRepo:    taskRepo,
		projectRepo: projectRepo,
	}
}

func (uc *taskUseCase) Create(ctx context.Context, input CreateTaskInput) (*repository.Task, error) {
	// Verifica se o projeto existe
	project, err := uc.projectRepo.FindByID(ctx, input.ProjectID)
	if err != nil {
		return nil, fmt.Errorf("erro ao verificar projeto: %v", err)
	}
	if project == nil {
		return nil, fmt.Errorf("projeto não encontrado")
	}

	task := &repository.Task{
		ID:          uuid.New().String(),
		ProjectID:   input.ProjectID,
		Title:       input.Title,
		Description: input.Description,
		Status:      repository.TaskStatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := uc.taskRepo.Create(ctx, *task); err != nil {
		return nil, fmt.Errorf("erro ao criar tarefa: %v", err)
	}

	return task, nil
}

func (uc *taskUseCase) GetByID(ctx context.Context, id string) (*repository.Task, error) {
	task, err := uc.taskRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar tarefa: %v", err)
	}
	if task == nil {
		return nil, fmt.Errorf("tarefa não encontrada")
	}
	return task, nil
}

func (uc *taskUseCase) List(ctx context.Context) ([]repository.Task, error) {
	tasks, err := uc.taskRepo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar tarefas: %v", err)
	}
	return tasks, nil
}

func (uc *taskUseCase) Update(ctx context.Context, input struct {
	ID          string
	Title       string
	Description string
}) (*repository.Task, error) {
	existingTask, err := uc.taskRepo.FindByID(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar tarefa: %v", err)
	}
	if existingTask == nil {
		return nil, fmt.Errorf("tarefa não encontrada")
	}

	task := &repository.Task{
		ID:          input.ID,
		ProjectID:   existingTask.ProjectID,
		Title:       input.Title,
		Description: input.Description,
		Status:      existingTask.Status,
		AssignedTo:  existingTask.AssignedTo,
		CreatedAt:   existingTask.CreatedAt,
		UpdatedAt:   time.Now(),
	}

	if err := uc.taskRepo.Update(ctx, *task); err != nil {
		return nil, fmt.Errorf("erro ao atualizar tarefa: %v", err)
	}

	return task, nil
}

func (uc *taskUseCase) UpdateStatus(ctx context.Context, input UpdateTaskStatusInput) error {
	existingTask, err := uc.taskRepo.FindByID(ctx, input.ID)
	if err != nil {
		return fmt.Errorf("erro ao buscar tarefa: %v", err)
	}
	if existingTask == nil {
		return fmt.Errorf("tarefa não encontrada")
	}

	existingTask.Status = input.Status
	existingTask.UpdatedAt = time.Now()

	if err := uc.taskRepo.Update(ctx, *existingTask); err != nil {
		return fmt.Errorf("erro ao atualizar status da tarefa: %v", err)
	}

	return nil
}

func (uc *taskUseCase) AssignTask(ctx context.Context, input AssignTaskInput) error {
	existingTask, err := uc.taskRepo.FindByID(ctx, input.TaskID)
	if err != nil {
		return fmt.Errorf("erro ao buscar tarefa: %v", err)
	}
	if existingTask == nil {
		return fmt.Errorf("tarefa não encontrada")
	}

	existingTask.AssignedTo = input.AssignedTo
	existingTask.UpdatedAt = time.Now()

	if err := uc.taskRepo.Update(ctx, *existingTask); err != nil {
		return fmt.Errorf("erro ao atribuir tarefa: %v", err)
	}

	return nil
}

func (uc *taskUseCase) ListByProject(ctx context.Context, projectID string) ([]repository.Task, error) {
	tasks, err := uc.taskRepo.FindByProject(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar tarefas do projeto: %v", err)
	}
	return tasks, nil
}

func (uc *taskUseCase) ListByAssignee(ctx context.Context, userID string) ([]repository.Task, error) {
	tasks, err := uc.taskRepo.FindByAssignee(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar tarefas do usuário: %v", err)
	}
	return tasks, nil
}
