package main

import (
	"context"
	"os"

	"github.com/indiependente/go-proverbs-server/pkg/logging"
	"github.com/indiependente/go-proverbs-server/server"
	"github.com/indiependente/pkg/logger"
	"github.com/indiependente/pkg/shutdown"
)

const (
	service = `proverber`
)

func main() {
	grpcPort := os.Getenv("GRPC_PORT")
	grpcHost := os.Getenv("GRPC_HOST")
	httpPort := os.Getenv("HTTP_PORT")

	log := logging.NewPLogger(service, logger.ParseLogLevel(os.Getenv("LOG_LEVEL")))

	// init server
	srv := server.New(
		log,
		server.Config{
			GRPCPort: grpcPort,
			GRPCHost: grpcHost,
			HTTPPort: httpPort,
		},
	)

	ctx, cancel := context.WithCancel(context.Background())
	// Start gRPC server
	go func() {
		err := srv.StartGRPC(ctx)
		if err != nil {
			log.Fatal("Error while running GRPC server", err)
		}
	}()

	// Start HTTP server
	go func() {
		err := srv.StartGW(ctx)
		if err != nil {
			log.Fatal("Error while running gRPC-gateway HTTP server", err)
		}
	}()

	// Wait
	err := shutdown.Wait(ctx, cancel, srv.Shutdown, log)
	if err != nil {
		log.Fatal("Error while shutting down server", err)
	}
}
