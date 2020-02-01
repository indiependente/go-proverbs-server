package server

import (
	"context"
	"net"
	"net/http"

	gw "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/indiependente/go-proverbs-server/pkg/logging"
	"github.com/indiependente/go-proverbs-server/rpc"
	grpctransport "github.com/indiependente/go-proverbs-server/transport/grpc"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Proverber is a server that responds with a proverb.
type Proverber interface {
	Proverb(context.Context) string
}

// Server implements a Proverber server.
type Server struct {
	grpcServer *grpc.Server
	httpServer *http.Server
	logger     logging.ProverbLogger
	config     Config
}

// Config holds server configurations.
type Config struct {
	GRPCPort string
	GRPCHost string
	HTTPPort string
}

// New returns a new instance of the Server.
func New(l logging.ProverbLogger, c Config) Server {
	return Server{
		logger: l,
		config: c,
	}
}

// StartGRPC starts the GRPC server.
func (srv *Server) StartGRPC(ctx context.Context) error {
	srv.logger.Event("start").Info("Starting GRPC server on port " + srv.config.GRPCPort)
	lis, err := net.Listen("tcp", ":"+srv.config.GRPCPort)
	if err != nil {
		return errors.Wrap(err, "Could not start grpc server")
	}

	srv.grpcServer = grpc.NewServer(withServerUnaryInterceptor())

	rpc.RegisterProverberServer(srv.grpcServer, srv)

	return srv.grpcServer.Serve(lis)
}

// StartGW starts the GRPC REST gateway.
func (srv *Server) StartGW(ctx context.Context) error {
	srv.logger.Event("start").Info("Starting HTTP server on port " + srv.config.HTTPPort)
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := gw.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := rpc.RegisterProverberHandlerFromEndpoint(ctx, mux, srv.config.GRPCHost+":"+srv.config.GRPCPort, opts)
	if err != nil {
		return errors.Wrap(err, "Could not start grpc-gateway")
	}

	srv.httpServer = &http.Server{
		Addr:    ":" + srv.config.HTTPPort,
		Handler: mux,
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	if err := srv.httpServer.ListenAndServe(); err != http.ErrServerClosed {
		srv.logger.Error("Failed to listen and serve", err)
		return errors.Wrap(err, "Failed to listen and serve")
	}
	return nil
}

// Shutdown initiates the graceful shutdown process.
func (srv *Server) Shutdown(ctx context.Context) error {
	defer srv.logger.Event("shutdown").Info("Exit")

	// Shut down HTTP server
	srv.logger.Event("shutdown").Info("Shutting down the http server")
	if srv.httpServer == nil {
		return errors.New("HTTP server is nil")
	}
	err := srv.httpServer.Shutdown(ctx)
	if err != nil {
		srv.logger.Error("Failed to shutdown http server", err)
		return errors.Wrap(err, "Failed to shutdown http server")
	}

	// Shut down GRPC server
	srv.logger.Event("shutdown").Info("Shutting down the grpc server")
	if srv.grpcServer == nil {
		return errors.New("GRPC server is nil")
	}
	srv.grpcServer.GracefulStop()

	// Add any graceful shutdown operations below, e.g. closing connections, etc.

	return nil
}

func withServerUnaryInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(grpctransport.ServerInterceptor)
}
