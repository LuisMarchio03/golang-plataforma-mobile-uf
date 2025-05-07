package handler

import (
	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/repository"
	usecase "github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/usecase/project_application"
	"github.com/gofiber/fiber/v2"
)

type ProjectApplicationHandler struct {
	useCase usecase.ProjectApplicationUseCase
}

func NewProjectApplicationHandler(useCase usecase.ProjectApplicationUseCase) *ProjectApplicationHandler {
	return &ProjectApplicationHandler{
		useCase: useCase,
	}
}

func (h *ProjectApplicationHandler) Create(c *fiber.Ctx) error {
	var input struct {
		ProjectID string `json:"project_id"`
		Message   string `json:"message"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Dados inválidos",
		})
	}

	// Obtém o ID do usuário do contexto (definido pelo middleware de autenticação)
	userID := c.Locals("userID").(string)

	application, err := h.useCase.Create(c.Context(), usecase.CreateProjectApplicationInput{
		ProjectID: input.ProjectID,
		UserID:    userID,
		Message:   input.Message,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(application)
}

func (h *ProjectApplicationHandler) ListByProject(c *fiber.Ctx) error {
	projectID := c.Params("projectId")
	if projectID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID do projeto não fornecido",
		})
	}

	applications, err := h.useCase.ListByProject(c.Context(), projectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(applications)
}

func (h *ProjectApplicationHandler) ListByUser(c *fiber.Ctx) error {
	userID := c.Params("userId")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID do usuário não fornecido",
		})
	}

	applications, err := h.useCase.ListByUser(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(applications)
}

func (h *ProjectApplicationHandler) UpdateStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID da candidatura não fornecido",
		})
	}

	var input struct {
		Status string `json:"status"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Dados inválidos",
		})
	}

	err := h.useCase.UpdateStatus(c.Context(), usecase.UpdateApplicationStatusInput{
		ID:     id,
		Status: repository.ApplicationStatus(input.Status),
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusOK)
}
