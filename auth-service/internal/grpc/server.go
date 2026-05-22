package grpc

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcServer struct {
	server *grpc.Server
	port   int
}

func NewServer(port int) *GrpcServer {
	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.PayloadReceived,
			logging.PayloadSent,
		),
	}

	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) error {
			slog.Error("recovered from panic", slog.Any("panic", p))
			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	server := grpc.NewServer(grpc.ChainUnaryInterceptor(
		logging.UnaryServerInterceptor(InterceptorLogger(), loggingOpts...),
		recovery.UnaryServerInterceptor(recoveryOpts...),
	))

	return &GrpcServer{
		server: server,
		port:   port,
	}
}

func InterceptorLogger() logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		slog.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func (s *GrpcServer) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("gprc run: %w", err)
	}

	log.Println("grpc server started", slog.String("addr", l.Addr().String()))
	if err := s.server.Serve(l); err != nil {
		return fmt.Errorf("grpc run err: %w", err)
	}

	return nil
}

func (s *GrpcServer) Stop() {
	log.Println("grpc server stopped")
	s.server.GracefulStop()
}
