package main

import (
	"fmt"
	"time"

	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/config"
	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/database"
	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/handler"
	customLogger "github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/logger"
	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/repository/postgres"
	"github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/router"
	projectUsecase "github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/usecase/project"
	projectAppUsecase "github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/usecase/project_application"
	taskUsecase "github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/usecase/task"
	userUsecase "github.com/LuisMarchio03/golang-plataforma-mobile-uf/internal/usecase/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Carrega configurações
	cfg, err := config.Load()
	if err != nil {
		fmt.Printf("Erro ao carregar configurações: %v\n", err)
		return
	}

	// Inicializa o logger
	log := customLogger.New(customLogger.INFO)

	// Configura a conexão com o banco de dados
	dbConfig := database.Config{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
	}

	db, err := database.NewConnection(dbConfig)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	log.Info("Conectado ao banco de dados com sucesso")

	// Inicializa os handlers
	// Initialize repositories
	userRepo := postgres.NewUserRepository(db)
	projectRepo := postgres.NewProjectRepository(db)
	projectApplicationRepo := postgres.NewProjectApplicationRepository(db)
	taskRepo := postgres.NewTaskRepository(db)

	// Initialize use cases
	userUseCase := userUsecase.NewUserUseCaseImpl(userRepo)
	projectUseCase := projectUsecase.NewProjectUseCase(projectRepo)
	projectApplicationUseCase := projectAppUsecase.NewProjectApplicationUseCase(projectApplicationRepo, projectRepo, userRepo)
	taskUseCase := taskUsecase.NewTaskUseCase(taskRepo, projectRepo)

	// Initialize handlers
	handlers := handler.NewHandler(
		handler.NewUserHandler(userUseCase),
		handler.NewProjectHandler(projectUseCase, projectApplicationUseCase),
		handler.NewProjectApplicationHandler(projectApplicationUseCase),
		handler.NewTaskHandler(taskUseCase),
	)

	// Inicia o servidor HTTP
	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	// Inicializa o Fiber
	app := fiber.New(fiber.Config{
		AppName:      "Plataforma Mobile API",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	})

	// Middlewares
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(fiberLogger.New(fiberLogger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))

	// Configura as rotas usando o pacote router
	router.SetupRoutes(app, handlers)

	// Inicia o servidor
	log.Info("Servidor iniciando na porta %s", cfg.ServerPort)
	if err := app.Listen(serverAddr); err != nil {
		log.Fatal("Erro ao iniciar o servidor: %v", err)
	}
}
