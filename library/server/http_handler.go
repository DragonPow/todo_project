package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type httpServer struct {
	server   *http.Server
	grpcConn *grpc.ClientConn
	config   *httpConfig
}

type httpConfig struct {
	Addr          ListenAddr
	Handlers      []HttpServerHandler
	MuxConfigs    []runtime.ServeMuxOption
	ServerConfigs []HttpServerConfig
}

type HttpServerConfig func(*http.Server)
type HttpServerHandler func(*http.ServeMux)

func newHttpServer(cfg *httpConfig, services []Service, grpcConn *grpc.ClientConn) (*httpServer, error) {
	mux := runtime.NewServeMux(cfg.MuxConfigs...)
	httpMux := http.NewServeMux()
	var err error

	// Handle http itself
	for _, handler := range cfg.Handlers {
		handler(httpMux)
	}

	// Register services
	for _, svc := range services {
		err = svc.RegisterHttp(context.Background(), mux, grpcConn)
		if err != nil {
			return nil, err
		}
	}

	// Handle grpc-gateway
	httpMux.Handle("/", mux)

	server := &http.Server{
		Addr:    cfg.Addr.String(),
		Handler: httpMux,
	}
	for _, c := range cfg.ServerConfigs {
		c(server)
	}

	return &httpServer{
		server:   server,
		config:   cfg,
		grpcConn: grpcConn,
	}, nil
}

func loadDefaultHttpConfig() *httpConfig {
	return &httpConfig{
		Addr: ListenAddr{
			Host: "0.0.0.0",
			Port: 10050,
		},
		MuxConfigs: []runtime.ServeMuxOption{
			// runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler),
		},
		Handlers:      []HttpServerHandler{},
		ServerConfigs: []HttpServerConfig{},
	}
}

func (s *httpServer) Start() error {
	s.grpcConn.Connect()
	fmt.Println("HTTP server started with address:", s.config.Addr.String())
	return s.server.ListenAndServe()
}

func (s *httpServer) Stop(ctx context.Context) error {
	s.grpcConn.Close()
	return s.server.Shutdown(ctx)
}
