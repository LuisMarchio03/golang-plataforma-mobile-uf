package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/repository"
)

type projectRepository struct {
	db *sql.DB
}

// Delete implements repository.ProjectRepository.
func (r *projectRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindByCreator implements repository.ProjectRepository.
func (r *projectRepository) FindByCreator(ctx context.Context, creatorID string) ([]repository.Project, error) {
	panic("unimplemented")
}

// FindByID implements repository.ProjectRepository.
func (r *projectRepository) FindByID(ctx context.Context, id string) (*repository.Project, error) {
	query := `
        SELECT id, title, description, status, created_by, created_at, updated_at
        FROM projects
        WHERE id = $1
    `

	var project repository.Project
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&project.ID,
		&project.Title,
		&project.Description,
		&project.Status,
		&project.CreatedBy,
		&project.CreatedAt,
		&project.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar projeto: %v", err)
	}

	return &project, nil
}

// FindByStatus implements repository.ProjectRepository.
func (r *projectRepository) FindByStatus(ctx context.Context, status repository.ProjectStatus) ([]repository.Project, error) {
	panic("unimplemented")
}

// List implements repository.ProjectRepository.
func (r *projectRepository) List(ctx context.Context) ([]repository.Project, error) {
	query := `
        SELECT id, title, description, status, created_by, created_at, updated_at
        FROM projects
    `

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar projetos: %v", err)
	}
	defer rows.Close()

	var projects []repository.Project
	for rows.Next() {
		var project repository.Project
		err := rows.Scan(
			&project.ID,
			&project.Title,
			&project.Description,
			&project.Status,
			&project.CreatedBy,
			&project.CreatedAt,
			&project.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler projeto: %v", err)
		}
		projects = append(projects, project)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao iterar projetos: %v", err)
	}

	return projects, nil
}

// Update implements repository.ProjectRepository.
func (r *projectRepository) Update(ctx context.Context, project repository.Project) error {
	query := `
        UPDATE projects
        SET title = $1,
            description = $2,
            status = $3,
            updated_at = $4
        WHERE id = $5
    `

	result, err := r.db.ExecContext(ctx, query,
		project.Title,
		project.Description,
		project.Status,
		project.UpdatedAt,
		project.ID,
	)

	if err != nil {
		return fmt.Errorf("erro ao atualizar projeto: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar atualização: %v", err)
	}

	if rows == 0 {
		return fmt.Errorf("projeto não encontrado")
	}

	return nil
}

func NewProjectRepository(db *sql.DB) repository.ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) Create(ctx context.Context, project repository.Project) error {
	query := `
		INSERT INTO projects (id, title, description, status, created_by, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.ExecContext(ctx, query,
		project.ID,
		project.Title,
		project.Description,
		project.Status,
		project.CreatedBy,
		project.CreatedAt,
		project.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("erro ao criar projeto: %v", err)
	}

	return nil
}
