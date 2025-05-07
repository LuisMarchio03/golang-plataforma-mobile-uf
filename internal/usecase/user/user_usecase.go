package usecase

import (
	"context"

	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/repository"
)

type UserUseCase interface {
	Create(ctx context.Context, input CreateUserInput) (*repository.User, error)
	GetByID(ctx context.Context, id string) (*repository.User, error)
	List(ctx context.Context) ([]repository.User, error)
	Update(ctx context.Context, input UpdateUserInput) (*repository.User, error)
	Delete(ctx context.Context, id string) error
	ValidateCredentials(ctx context.Context, email, password string) (*repository.User, error)
}

type CreateUserInput struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UpdateUserInput struct {
	ID       string `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"omitempty,min=6"`
}

func NewUserUseCaseImpl(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}
