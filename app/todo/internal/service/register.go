package service

import (
	"context"

	"github.com/DragonPow/todo_project/app/todo/api"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func (s *Service) RegisterGrpc(server *grpc.Server) {
	api.RegisterTodoServiceServer(server, s)
}

func (s *Service) RegisterHttp(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return api.RegisterTodoServiceHandler(ctx, mux, conn)
}
