package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/repository"
)

type taskRepository struct {
	db *sql.DB
}

// Delete implements repository.TaskRepository.
func (r *taskRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindByAssignee implements repository.TaskRepository.
func (r *taskRepository) FindByAssignee(ctx context.Context, assigneeID string) ([]repository.Task, error) {
	panic("unimplemented")
}

// FindByStatus implements repository.TaskRepository.
func (r *taskRepository) FindByStatus(ctx context.Context, status repository.TaskStatus) ([]repository.Task, error) {
	panic("unimplemented")
}

// List implements repository.TaskRepository.
func (r *taskRepository) List(ctx context.Context) ([]repository.Task, error) {
	panic("unimplemented")
}

func NewTaskRepository(db *sql.DB) repository.TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) Create(ctx context.Context, task repository.Task) error {
	query := `
        INSERT INTO tasks (id, project_id, title, description, status, assigned_to, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `

	_, err := r.db.ExecContext(ctx, query,
		task.ID,
		task.ProjectID,
		task.Title,
		task.Description,
		task.Status,
		task.AssignedTo,
		task.CreatedAt,
		task.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("erro ao criar tarefa: %v", err)
	}

	return nil
}

func (r *taskRepository) FindByID(ctx context.Context, id string) (*repository.Task, error) {
	query := `
        SELECT id, project_id, title, description, status, assigned_to, created_at, updated_at
        FROM tasks
        WHERE id = $1
    `

	var task repository.Task
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&task.ID,
		&task.ProjectID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.AssignedTo,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar tarefa: %v", err)
	}

	return &task, nil
}

func (r *taskRepository) FindByProject(ctx context.Context, projectID string) ([]repository.Task, error) {
	query := `
        SELECT id, project_id, title, description, status, assigned_to, created_at, updated_at
        FROM tasks
        WHERE project_id = $1
    `

	rows, err := r.db.QueryContext(ctx, query, projectID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar tarefas do projeto: %v", err)
	}
	defer rows.Close()

	var tasks []repository.Task
	for rows.Next() {
		var task repository.Task
		err := rows.Scan(
			&task.ID,
			&task.ProjectID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.AssignedTo,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler tarefa: %v", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *taskRepository) Update(ctx context.Context, task repository.Task) error {
	query := `
        UPDATE tasks
        SET title = $1, description = $2, status = $3, assigned_to = $4, updated_at = $5
        WHERE id = $6
    `

	result, err := r.db.ExecContext(ctx, query,
		task.Title,
		task.Description,
		task.Status,
		task.AssignedTo,
		task.UpdatedAt,
		task.ID,
	)

	if err != nil {
		return fmt.Errorf("erro ao atualizar tarefa: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar atualização: %v", err)
	}

	if rows == 0 {
		return fmt.Errorf("tarefa não encontrada")
	}

	return nil
}
