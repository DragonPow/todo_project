package log

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnaryServerPanicInterceptor(logger *Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				logger.Error("Function is panic", zap.Any("panic", r))
				err = status.Error(codes.Internal, "Fail to process request")
				return
			}
		}()
		resp, err = handler(ctx, req)
		return resp, err
	}
}
