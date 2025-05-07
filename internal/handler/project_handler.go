package handler

import (
	usecase "github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/usecase/project"
	"github.com/gofiber/fiber/v2"
)

type ProjectHandler struct {
	projectUseCase usecase.ProjectUseCase
}

func NewProjectHandler(projectUseCase usecase.ProjectUseCase) *ProjectHandler {
	return &ProjectHandler{
		projectUseCase: projectUseCase,
	}
}

func (h *ProjectHandler) Create(c *fiber.Ctx) error {
	var input usecase.CreateProjectInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(NewErrorResponse("erro ao processar dados"))
	}

	project, err := h.projectUseCase.Create(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(NewSuccessResponse(project))
}

func (h *ProjectHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	project, err := h.projectUseCase.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(NewErrorResponse(err.Error()))
	}

	return c.JSON(NewSuccessResponse(project))
}

func (h *ProjectHandler) List(c *fiber.Ctx) error {
	// h.logger.Info("Iniciando listagem de projetos")

	// Obtém o ID do usuário do contexto (setado pelo middleware de auth)
	userID := c.Locals("userID").(string)
	if userID == "" {
		// h.logger.Error("ID do usuário não encontrado no contexto")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Usuário não autenticado",
		})
	}

	projects, err := h.projectUseCase.List(c.Context())
	if err != nil {
		// h.logger.Error("Erro ao listar projetos: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao listar projetos",
		})
	}

	// h.logger.Info("Projetos listados com sucesso")
	return c.Status(fiber.StatusOK).JSON(projects)
}

func (h *ProjectHandler) Update(c *fiber.Ctx) error {
	var input usecase.UpdateProjectInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(NewErrorResponse("erro ao processar dados"))
	}

	input.ID = c.Params("id")
	project, err := h.projectUseCase.Update(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse(err.Error()))
	}

	return c.JSON(NewSuccessResponse(project))
}

func (h *ProjectHandler) UpdateStatus(c *fiber.Ctx) error {
	var input usecase.UpdateProjectStatusInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(NewErrorResponse("erro ao processar dados"))
	}

	input.ID = c.Params("id")
	if err := h.projectUseCase.UpdateStatus(c.Context(), input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(NewSuccessResponse(nil))
}
