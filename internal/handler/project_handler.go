package handler

import (
	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/repository"
	usecase "github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/usecase/project"
	projectApplicationUseCase "github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/usecase/project_application"
	"github.com/gofiber/fiber/v2"
)

type ProjectHandler struct {
	projectUseCase            usecase.ProjectUseCase
	projectApplicationUseCase projectApplicationUseCase.ProjectApplicationUseCase
}

func NewProjectHandler(projectUseCase usecase.ProjectUseCase, projectApplicationUseCase projectApplicationUseCase.ProjectApplicationUseCase) *ProjectHandler {
	return &ProjectHandler{
		projectUseCase:            projectUseCase,
		projectApplicationUseCase: projectApplicationUseCase,
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

func (h *ProjectHandler) ListEnrolled(c *fiber.Ctx) error {
	// Get current user ID from context (set by auth middleware)
	userID := c.Locals("userID").(string)

	// Get all applications for this user
	applications, err := h.projectApplicationUseCase.ListByUser(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch user applications",
		})
	}

	// Get all projects where user has an approved application
	var enrolledProjects []*repository.Project // Changed to slice of pointers
	for _, app := range applications {
		if app.Status == "approved" || app.Status == "pending" {
			project, err := h.projectUseCase.GetByID(c.Context(), app.ProjectID)
			if err != nil {
				continue // Skip if project not found
			}
			enrolledProjects = append(enrolledProjects, project) // Now this works because types match
		}
	}

	return c.JSON(fiber.Map{
		"projects": enrolledProjects,
	})
}
