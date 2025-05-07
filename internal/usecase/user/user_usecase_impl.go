package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo repository.UserRepository
}

// ValidateCredentials implements UserUseCase.
func (uc *userUseCase) ValidateCredentials(ctx context.Context, email string, password string) (*repository.User, error) {
    user, err := uc.userRepo.FindByEmail(ctx, email)
    if err != nil {
        return nil, fmt.Errorf("erro ao buscar usuário: %v", err)
    }
    if user == nil {
        return nil, fmt.Errorf("credenciais inválidas")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return nil, fmt.Errorf("credenciais inválidas")
    }

    return user, nil
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) Create(ctx context.Context, input CreateUserInput) (*repository.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("erro ao criptografar senha: %v", err)
	}

	user := &repository.User{
		ID:        uuid.New().String(),
		Name:      input.Name,
		Email:     input.Email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.userRepo.Create(ctx, *user); err != nil {
		return nil, fmt.Errorf("erro ao criar usuário: %v", err)
	}

	return user, nil
}

func (uc *userUseCase) GetByID(ctx context.Context, id string) (*repository.User, error) {
	user, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuário: %v", err)
	}
	if user == nil {
		return nil, fmt.Errorf("usuário não encontrado")
	}
	return user, nil
}

func (uc *userUseCase) List(ctx context.Context) ([]repository.User, error) {
	users, err := uc.userRepo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar usuários: %v", err)
	}
	return users, nil
}

func (uc *userUseCase) Update(ctx context.Context, input UpdateUserInput) (*repository.User, error) {
	// Verifica se o usuário existe
	existingUser, err := uc.userRepo.FindByID(ctx, input.ID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar usuário: %v", err)
	}
	if existingUser == nil {
		return nil, fmt.Errorf("usuário não encontrado")
	}

	// Verifica se o novo email já está em uso por outro usuário
	if input.Email != existingUser.Email {
		userWithEmail, err := uc.userRepo.FindByEmail(ctx, input.Email)
		if err != nil {
			return nil, fmt.Errorf("erro ao verificar email existente: %v", err)
		}
		if userWithEmail != nil && userWithEmail.ID != input.ID {
			return nil, fmt.Errorf("email já está em uso por outro usuário")
		}
	}

	// Atualiza os dados do usuário
	user := &repository.User{
		ID:        input.ID,
		Name:      input.Name,
		Email:     input.Email,
		Password:  existingUser.Password,
		CreatedAt: existingUser.CreatedAt,
		UpdatedAt: time.Now(),
	}

	// Se uma nova senha foi fornecida, atualiza a senha
	if input.Password != "" {
		user.Password = input.Password // Aqui você deve adicionar hash da senha
	}

	// Salva as alterações
	if err := uc.userRepo.Update(ctx, *user); err != nil {
		return nil, fmt.Errorf("erro ao atualizar usuário: %v", err)
	}

	return user, nil
}

func (uc *userUseCase) Delete(ctx context.Context, id string) error {
	// Verifica se o usuário existe
	existingUser, err := uc.userRepo.FindByID(ctx, id)
	if err != nil {
		return fmt.Errorf("erro ao buscar usuário: %v", err)
	}
	if existingUser == nil {
		return fmt.Errorf("usuário não encontrado")
	}

	// Deleta o usuário
	if err := uc.userRepo.Delete(ctx, id); err != nil {
		return fmt.Errorf("erro ao deletar usuário: %v", err)
	}

	return nil
}
