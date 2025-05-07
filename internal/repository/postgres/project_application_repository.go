package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/repository"
)

type projectApplicationRepository struct {
	db *sql.DB
}

// Delete implements repository.ProjectApplicationRepository.
func (r *projectApplicationRepository) Delete(ctx context.Context, id string) error {
	panic("unimplemented")
}

// FindByStatus implements repository.ProjectApplicationRepository.
func (r *projectApplicationRepository) FindByStatus(ctx context.Context, status repository.ApplicationStatus) ([]repository.ProjectApplication, error) {
	panic("unimplemented")
}

// FindByUser implements repository.ProjectApplicationRepository.
func (r *projectApplicationRepository) FindByUser(ctx context.Context, userID string) ([]repository.ProjectApplication, error) {
	query := `
        SELECT id, project_id, user_id, message, status, created_at, updated_at
        FROM project_applications
        WHERE user_id = $1
    `

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar candidaturas do usuário: %v", err)
	}
	defer rows.Close()

	var applications []repository.ProjectApplication
	for rows.Next() {
		var app repository.ProjectApplication
		err := rows.Scan(
			&app.ID,
			&app.ProjectID,
			&app.UserID,
			&app.Message,
			&app.Status,
			&app.CreatedAt,
			&app.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler candidatura: %v", err)
		}
		applications = append(applications, app)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao iterar candidaturas: %v", err)
	}

	return applications, nil
}

// List implements repository.ProjectApplicationRepository.
func (r *projectApplicationRepository) List(ctx context.Context) ([]repository.ProjectApplication, error) {
	panic("unimplemented")
}

func NewProjectApplicationRepository(db *sql.DB) repository.ProjectApplicationRepository {
	return &projectApplicationRepository{
		db: db,
	}
}

func (r *projectApplicationRepository) Create(ctx context.Context, application repository.ProjectApplication) error {
	query := `
        INSERT INTO project_applications (id, project_id, user_id, message, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `

	_, err := r.db.ExecContext(ctx, query,
		application.ID,
		application.ProjectID,
		application.UserID,
		application.Message,
		application.Status,
		application.CreatedAt,
		application.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("erro ao criar candidatura: %v", err)
	}

	return nil
}

func (r *projectApplicationRepository) FindByID(ctx context.Context, id string) (*repository.ProjectApplication, error) {
	query := `
        SELECT id, project_id, user_id, message, status, created_at, updated_at
        FROM project_applications
        WHERE id = $1
    `

	var app repository.ProjectApplication
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&app.ID,
		&app.ProjectID,
		&app.UserID,
		&app.Message,
		&app.Status,
		&app.CreatedAt,
		&app.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar candidatura: %v", err)
	}

	return &app, nil
}

func (r *projectApplicationRepository) FindByProject(ctx context.Context, projectID string) ([]repository.ProjectApplication, error) {
	query := `
        SELECT id, project_id, user_id, message, status, created_at, updated_at
        FROM project_applications
        WHERE project_id = $1
    `

	rows, err := r.db.QueryContext(ctx, query, projectID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar candidaturas do projeto: %v", err)
	}
	defer rows.Close()

	var applications []repository.ProjectApplication
	for rows.Next() {
		var app repository.ProjectApplication
		err := rows.Scan(
			&app.ID,
			&app.ProjectID,
			&app.UserID,
			&app.Message,
			&app.Status,
			&app.CreatedAt,
			&app.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler candidatura: %v", err)
		}
		applications = append(applications, app)
	}

	return applications, nil
}

func (r *projectApplicationRepository) Update(ctx context.Context, application repository.ProjectApplication) error {
	query := `
        UPDATE project_applications
        SET status = $1, updated_at = $2
        WHERE id = $3
    `

	result, err := r.db.ExecContext(ctx, query,
		application.Status,
		application.UpdatedAt,
		application.ID,
	)

	if err != nil {
		return fmt.Errorf("erro ao atualizar candidatura: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar atualização: %v", err)
	}

	if rows == 0 {
		return fmt.Errorf("candidatura não encontrada")
	}

	return nil
}
