package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ix1ax/social-network-backend/internal/common/config"
	"github.com/ix1ax/social-network-backend/internal/common/db"
	"github.com/ix1ax/social-network-backend/internal/user/handler"
	"github.com/ix1ax/social-network-backend/internal/user/repository"
	"github.com/ix1ax/social-network-backend/internal/user/service"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type App struct {
	cfg    *config.Config
	db     *sqlx.DB
	router *gin.Engine
}

func New() (*App, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, fmt.Errorf("error loading env files: %w", err)
	}

	cfg := config.Load()

	database, err := db.NewConnection(cfg)

	if err != nil {
		return nil, fmt.Errorf("error connectiong to database: %w", err)
	}

	if err := db.RunMigrations(cfg); err != nil {
		return nil, fmt.Errorf("error running migrations: %w", err)
	}

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	v1 := router.Group("/api/v1")
	userHandler.RegisterRoutes(v1)

	return &App{
		cfg:    cfg,
		db:     database,
		router: router,
	}, nil
}

func (a *App) Run() error {
	/*
	 * defer откладывает выполнение функции до момента выхода из текущей
	 * т.е как только функция main() завершится, закроется и подключение к БД
	 */
	defer a.db.Close()
	return a.router.Run(":" + a.cfg.ServerPort)
}
