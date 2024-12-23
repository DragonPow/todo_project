package server

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	server *grpc.Server
	config *grpcConfig
}

type grpcConfig struct {
	Addr          ListenAddr
	ServerConfigs []grpcServerConfig
}

type grpcServerConfig grpc.ServerOption

func newGrpcServer(cfg *grpcConfig, services []Service) (*grpcServer, error) {
	opts := []grpc.ServerOption{}
	for _, c := range cfg.ServerConfigs {
		opts = append(opts, grpc.ServerOption(c))
	}
	server := grpc.NewServer(opts...)
	reflection.Register(server)

	// Register services
	for _, svc := range services {
		svc.RegisterGrpc(server)
	}

	return &grpcServer{
		server: server,
		config: cfg,
	}, nil
}

func loadDefaultGrpcConfig() *grpcConfig {
	return &grpcConfig{
		Addr: ListenAddr{
			Host: "0.0.0.0",
			Port: 10051,
		},
	}
}

func (s *grpcServer) Start() error {
	lis, err := net.Listen("tcp", s.config.Addr.String())
	if err != nil {
		return err
	}
	fmt.Println("GRPC server started with address:", s.config.Addr.String())
	return s.server.Serve(lis)
}

func (s *grpcServer) Stop(_ context.Context) error {
	s.server.GracefulStop()
	return nil
}
