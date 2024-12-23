package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/DragonPow/todo_project/app/todo/internal/model"
)

type Repository interface {
	Close() error

	CreateTodo(ctx context.Context, req *model.Todo) (int, error)
	GetTodoByID(ctx context.Context, id int) (*model.Todo, error)
	SearchTodoByTitle(ctx context.Context, title string) ([]*model.Todo, error)
	UpdateTodoByID(ctx context.Context, id int, req *model.Todo) error
	DeleteTodoByID(ctx context.Context, id int) error
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) (Repository, error) {
	return &repo{
		db: db,
	}, nil
}

func (r *repo) Close() error {
	db, err := r.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}
