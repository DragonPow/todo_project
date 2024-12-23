package service

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/DragonPow/todo_project/app/todo/api"
	"github.com/DragonPow/todo_project/app/todo/config"
	"github.com/DragonPow/todo_project/app/todo/internal/repository"
	"github.com/DragonPow/todo_project/library/log"
)

type Service struct {
	cfg    *config.Config
	logger *log.Logger
	repo   repository.Repository

	api.UnimplementedTodoServiceServer
}

func NewService(cfg *config.Config, logger *log.Logger) (*Service, error) {
	logger.Debug("Database dsn", zap.String("dsn", cfg.Database.DSN()))
	db, err := gorm.Open(postgres.Open(cfg.Database.DSN()), cfg.GormConfig())
	if err != nil {
		fmt.Println("Error when open database", err)
		// return nil, err
	} else {
		logger.Info("Database connected")
		// Ping database
		if sqlDB, err := db.DB(); err == nil {
			if err := sqlDB.Ping(); err == nil {
				logger.Debug("Database pinged")
			}
		}
	}

	todoRepo, err := repository.NewRepository(db)
	if err != nil {
		return nil, err
	}

	return &Service{
		cfg:    cfg,
		logger: logger,
		repo:   todoRepo,
	}, nil
}

func (s *Service) Stop(_ context.Context) error {
	return s.repo.Close()
}
