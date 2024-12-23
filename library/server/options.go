package server

import (
	"google.golang.org/grpc"
)

type Option func(*Config)

func WithGrpcListener(lis ListenAddr) Option {
	return func(c *Config) {
		c.grpcConfig.Addr = ListenAddr{
			Host: lis.Host,
			Port: lis.Port,
		}
	}
}

func WithHttpListener(lis ListenAddr) Option {
	return func(c *Config) {
		c.httpConfig.Addr = ListenAddr{
			Host: lis.Host,
			Port: lis.Port,
		}
	}
}

func WithService(services ...Service) Option {
	return func(c *Config) {
		c.services = append(c.services, services...)
	}
}

func GrpcUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) Option {
	return func(c *Config) {
		var opts []grpcServerConfig
		for _, i := range interceptors {
			opts = append(opts, grpc.ChainUnaryInterceptor([]grpc.UnaryServerInterceptor{i}...))
		}
		c.grpcConfig.ServerConfigs = append(c.grpcConfig.ServerConfigs, opts...)
	}
}
