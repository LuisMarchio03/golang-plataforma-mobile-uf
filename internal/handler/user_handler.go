package handler

import (
	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/auth"
	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/usecase/user"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var input usecase.CreateUserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(NewErrorResponse("erro ao processar dados"))
	}

	user, err := h.userUseCase.Create(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(NewSuccessResponse(user))
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userUseCase.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(NewErrorResponse(err.Error()))
	}

	return c.JSON(NewSuccessResponse(user))
}

func (h *UserHandler) List(c *fiber.Ctx) error {
	users, err := h.userUseCase.List(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse(err.Error()))
	}

	return c.JSON(NewSuccessResponse(users))
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	var input usecase.UpdateUserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(NewErrorResponse("erro ao processar dados"))
	}

	input.ID = c.Params("id")
	user, err := h.userUseCase.Update(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse(err.Error()))
	}

	return c.JSON(NewSuccessResponse(user))
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.userUseCase.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(NewErrorResponse("erro ao processar dados"))
	}

	user, err := h.userUseCase.ValidateCredentials(c.Context(), input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(NewErrorResponse("credenciais inválidas"))
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse("erro ao gerar token"))
	}

	return c.JSON(NewSuccessResponse(fiber.Map{
		"token": token,
		"user":  user,
	}))
}