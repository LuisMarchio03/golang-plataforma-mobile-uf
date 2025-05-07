package repository

import "context"

// Repository interface base para operações comuns
type Repository[T any] interface {
    Create(ctx context.Context, entity T) error
    FindByID(ctx context.Context, id string) (*T, error)
    Update(ctx context.Context, entity T) error
    Delete(ctx context.Context, id string) error
    List(ctx context.Context) ([]T, error)
}