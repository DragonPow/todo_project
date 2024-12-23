package service

import (
	"context"
	"database/sql"

	"go.uber.org/zap"
	"google.golang.org/genproto/googleapis/rpc/code"

	"github.com/DragonPow/todo_project/app/todo/api"
	"github.com/DragonPow/todo_project/app/todo/internal/model"
)

func (s *Service) CreateTodo(ctx context.Context, req *api.CreateTodoRequest) (*api.CreateTodoResponse, error) {
	newTodo := &model.Todo{}
	newTodo.FromApiProto(req.Todo)
	todoID, err := s.repo.CreateTodo(ctx, newTodo)
	if err != nil {
		s.logger.Error("CreateTodo fail", zap.Error(err))
		return nil, err
	}

	todo, err := s.repo.GetTodoByID(ctx, todoID)
	if err != nil {
		s.logger.Error("GetTodoByID fail", zap.Error(err))
		return nil, err
	}

	return &api.CreateTodoResponse{
		Code:    code.Code_OK,
		Message: "OK",
		Todo:    todo.ToApiProto(),
	}, nil
}

func (s *Service) GetTodo(ctx context.Context, req *api.GetTodoRequest) (*api.GetTodoResponse, error) {
	todo, err := s.repo.GetTodoByID(ctx, int(req.Id))
	if err != nil {
		s.logger.Error("GetTodoByID fail", zap.Error(err))
		return nil, err
	}
	return &api.GetTodoResponse{
		Code:    code.Code_OK,
		Message: "OK",
		Todo:    todo.ToApiProto(),
	}, nil
}

func (s *Service) ListTodos(ctx context.Context, req *api.ListTodosRequest) (*api.ListTodosResponse, error) {
	todos, err := s.repo.SearchTodoByTitle(ctx, req.LikeTitle)
	if err != nil {
		s.logger.Error("SearchTodoByTitle fail", zap.Error(err))
		return nil, err
	}
	response := make([]*api.Todo, len(todos))
	for i, todo := range todos {
		response[i] = todo.ToApiProto()
		response[i].Items = []*api.Item{} // List items is not required
	}
	return &api.ListTodosResponse{
		Code:    code.Code_OK,
		Message: "OK",
		Todos:   response,
	}, nil
}

func (s *Service) UpdateTodo(ctx context.Context, req *api.UpdateTodoRequest) (*api.UpdateTodoResponse, error) {
	err := s.repo.UpdateTodoByID(ctx, int(req.Id), &model.Todo{
		Title:       req.Title,
		Description: sql.NullString{String: req.Description, Valid: req.Description != ""},
		IsDone:      req.Completed,
	})
	if err != nil {
		s.logger.Error("UpdateTodoByID fail", zap.Error(err))
		return nil, err
	}
	return &api.UpdateTodoResponse{
		Code:    code.Code_OK,
		Message: "OK",
	}, nil
}

func (s *Service) DeleteTodo(ctx context.Context, req *api.DeleteTodoRequest) (*api.DeleteTodoResponse, error) {
	err := s.repo.DeleteTodoByID(ctx, int(req.Id))
	if err != nil {
		s.logger.Error("DeleteTodoByID fail", zap.Error(err))
		return nil, err
	}
	return &api.DeleteTodoResponse{
		Code:    code.Code_OK,
		Message: "OK",
	}, nil
}
