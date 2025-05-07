package handler

import (
	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/usecase/task"
	"github.com/gofiber/fiber/v2"
)

type TaskHandler struct {
	taskUseCase usecase.TaskUseCase
}

func NewTaskHandler(taskUseCase usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{
		taskUseCase: taskUseCase,
	}
}

func (h *TaskHandler) Create(c *fiber.Ctx) error {
	var input usecase.CreateTaskInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(NewErrorResponse("erro ao processar dados"))
	}

	task, err := h.taskUseCase.Create(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(NewSuccessResponse(task))
}

func (h *TaskHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	task, err := h.taskUseCase.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(NewErrorResponse(err.Error()))
	}

	return c.JSON(NewSuccessResponse(task))
}

func (h *TaskHandler) ListByProject(c *fiber.Ctx) error {
	projectID := c.Params("projectId")
	tasks, err := h.taskUseCase.ListByProject(c.Context(), projectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse(err.Error()))
	}

	return c.JSON(NewSuccessResponse(tasks))
}

func (h *TaskHandler) UpdateStatus(c *fiber.Ctx) error {
	var input usecase.UpdateTaskStatusInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(NewErrorResponse("erro ao processar dados"))
	}

	input.ID = c.Params("id")
	if err := h.taskUseCase.UpdateStatus(c.Context(), input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(NewSuccessResponse(nil))
}

func (h *TaskHandler) AssignTask(c *fiber.Ctx) error {
	var input usecase.AssignTaskInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(NewErrorResponse("erro ao processar dados"))
	}

	input.TaskID = c.Params("id")
	if err := h.taskUseCase.AssignTask(c.Context(), input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(NewErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(NewSuccessResponse(nil))
}