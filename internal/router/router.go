package router

import (
	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/handler"
	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, handlers *handler.Handler) {
	// Rotas públicas
	api := app.Group("/api")

	// Rotas de autenticação
	auth := api.Group("/auth")
	auth.Post("/login", handlers.User.Login)
	auth.Post("/register", handlers.User.Create)

	// Rotas protegidas
	protected := api.Group("/", middleware.AuthMiddleware())

	// Rotas de usuário
	users := protected.Group("/users")
	users.Get("/", handlers.User.List)
	users.Get("/:id", handlers.User.GetByID)
	users.Put("/:id", handlers.User.Update)
	users.Delete("/:id", handlers.User.Delete)

	// Rotas de projeto
	projects := protected.Group("/projects")
	projects.Post("/", handlers.Project.Create)
	projects.Get("/", handlers.Project.List)
	projects.Get("/:id", handlers.Project.GetByID)
	projects.Put("/:id", handlers.Project.Update)
	projects.Put("/:id/status", handlers.Project.UpdateStatus)

	// Rotas de candidatura
	applications := protected.Group("/applications")
	applications.Post("/", handlers.ProjectApplication.Create)
	applications.Get("/project/:projectId", handlers.ProjectApplication.ListByProject)
	applications.Get("/user/:userId", handlers.ProjectApplication.ListByUser)
	applications.Put("/:id/status", handlers.ProjectApplication.UpdateStatus)

	// Rotas de tarefas
	tasks := protected.Group("/tasks")
	tasks.Post("/", handlers.Task.Create)
	tasks.Get("/project/:projectId", handlers.Task.ListByProject)
	tasks.Put("/:id/status", handlers.Task.UpdateStatus)
	tasks.Put("/:id/assign", handlers.Task.AssignTask)
}
