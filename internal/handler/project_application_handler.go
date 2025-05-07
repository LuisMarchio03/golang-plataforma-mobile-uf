package handler

import (
	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/usecase/project_application"
	"github.com/gofiber/fiber/v2"
)

type ProjectApplicationHandler struct {
	applicationUseCase usecase.ProjectApplicationUseCase
}

func NewProjectApplicationHandler(applicationUseCase usecase.ProjectApplicationUseCase) *ProjectApplicationHandler {
	return &ProjectApplicationHandler{
		applicationUseCase: applicationUseCase,
	}
}

func (h *ProjectApplicationHandler) Create(c *fiber.Ctx) error {
	var input usecase.CreateProjectApplicationInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(NewErrorResponse("erro ao processar dados"))
	}

	application, err := h.applicationUseCase.Create(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(NewSuccessResponse(application))
}

func (h *ProjectApplicationHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	application, err := h.applicationUseCase.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(NewErrorResponse(err.Error()))
	}

	return c.JSON(NewSuccessResponse(application))
}

func (h *ProjectApplicationHandler) ListByProject(c *fiber.Ctx) error {
	projectID := c.Params("projectId")
	applications, err := h.applicationUseCase.ListByProject(c.Context(), projectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse(err.Error()))
	}

	return c.JSON(NewSuccessResponse(applications))
}

func (h *ProjectApplicationHandler) ListByUser(c *fiber.Ctx) error {
	userID := c.Params("userId")
	applications, err := h.applicationUseCase.ListByUser(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse(err.Error()))
	}

	return c.JSON(NewSuccessResponse(applications))
}

func (h *ProjectApplicationHandler) UpdateStatus(c *fiber.Ctx) error {
	var input usecase.UpdateApplicationStatusInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(NewErrorResponse("erro ao processar dados"))
	}

	input.ID = c.Params("id")
	if err := h.applicationUseCase.UpdateStatus(c.Context(), input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(NewSuccessResponse(nil))
}