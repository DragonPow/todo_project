package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	GrpcServer *grpcServer
	HttpServer *httpServer
	services   []Service
}

func NewServer(opts ...Option) (*Server, error) {
	cfg := loadDefaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	log.Println("Start gprc server")
	grpcServer, err := newGrpcServer(cfg.grpcConfig, cfg.services)
	if err != nil {
		log.Println("Error creating grpc server:", err)
		return nil, err
	}

	log.Println("Start http server")
	grpcConn, err := grpc.NewClient(
		cfg.grpcConfig.Addr.String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println("Error creating grpc client:", err)
		return nil, err
	}
	// defer grpcConn.Close()
	log.Println("GRPC client connecting, current state:", grpcConn.GetState())

	httpServer, err := newHttpServer(cfg.httpConfig, cfg.services, grpcConn)
	if err != nil {
		log.Println("Error creating http server:", err)
		return nil, err
	}

	return &Server{
		GrpcServer: grpcServer,
		HttpServer: httpServer,
		services:   cfg.services,
	}, nil
}

func (s *Server) Start() error {
	errCh := make(chan error, 1)
	signalEndServices := []os.Signal{os.Interrupt, os.Kill}
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, signalEndServices...)

	go func() {
		errCh <- s.GrpcServer.Start()
	}()

	go func() {
		errCh <- s.HttpServer.Start()
	}()

	for {
		select {
		case <-stopCh:
			log.Println("Stopping server...")
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			for _, svc := range s.services {
				svc.Stop(ctx)
			}

			s.GrpcServer.Stop(ctx)
			s.HttpServer.Stop(ctx)
			return nil
		case err := <-errCh:
			log.Printf("Server error: %s\n", err)
			return err
		}
	}
}
