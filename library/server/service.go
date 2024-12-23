package server

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Service interface {
	RegisterGrpc(server *grpc.Server)
	RegisterHttp(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
	Stop(ctx context.Context) error
}
