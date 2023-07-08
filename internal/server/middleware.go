package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewMiddlewareConfig(logger log.Logger) []middleware.Middleware {
	return []middleware.Middleware{
		recovery.Recovery(),
		metadata.Server(),
		validate.Validator(),
		// tracing.Server(),
		// logging.Server(logger),
	}
}

// NewHTTPMiddleware
func NewHTTPMiddleware(logger log.Logger) http.ServerOption {
	return http.Middleware(
		NewMiddlewareConfig(logger)...,
	)
}

// NewGRPCMiddleware
func NewGRPCMiddleware(logger log.Logger) grpc.ServerOption {
	return grpc.Middleware(
		NewMiddlewareConfig(logger)...,
	)
}
