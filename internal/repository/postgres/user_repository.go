package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/repository"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user repository.User) error {
	query := `
        INSERT INTO users (id, name, email, password, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6)
    `

	now := time.Now()
	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		now,
		now,
	)

	if err != nil {
		return fmt.Errorf("erro ao criar usuário: %v", err)
	}

	return nil
}

func (r *userRepository) FindByID(ctx context.Context, id string) (*repository.User, error) {
	query := `
        SELECT id, name, email, password, created_at, updated_at
        FROM users
        WHERE id = $1
    `

	var user repository.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuário: %v", err)
	}

	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user repository.User) error {
	query := `
        UPDATE users
        SET name = $1, email = $2, password = $3, updated_at = $4
        WHERE id = $5
    `

	result, err := r.db.ExecContext(ctx, query,
		user.Name,
		user.Email,
		user.Password,
		time.Now(),
		user.ID,
	)

	if err != nil {
		return fmt.Errorf("erro ao atualizar usuário: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar linhas afetadas: %v", err)
	}

	if rows == 0 {
		return fmt.Errorf("nenhum usuário encontrado com o ID %s", user.ID)
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("erro ao deletar usuário: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("erro ao verificar linhas afetadas: %v", err)
	}

	if rows == 0 {
		return fmt.Errorf("nenhum usuário encontrado com o ID %s", id)
	}

	return nil
}

func (r *userRepository) List(ctx context.Context) ([]repository.User, error) {
	query := `
        SELECT id, name, email, password, created_at, updated_at
        FROM users
    `

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar usuários: %v", err)
	}
	defer rows.Close()

	var users []repository.User
	for rows.Next() {
		var user repository.User
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler usuário: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*repository.User, error) {
	query := `
        SELECT id, name, email, password, created_at, updated_at
        FROM users
        WHERE email = $1
    `

	var user repository.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuário por email: %v", err)
	}

	return &user, nil
}
