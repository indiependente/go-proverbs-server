package main

import (
	"context"
	"log"

	"github.com/indiependente/go-proverbs-server/pkg/logging"
	"github.com/indiependente/go-proverbs-server/server"
	"github.com/indiependente/pkg/logger"
	"github.com/indiependente/pkg/shutdown"
)

const (
	service = `proverber`
)

func main() {
	conf, err := parseConfig()
	if err != nil {
		log.Fatal(err)
	}
	log := logging.NewPLogger(service, logger.ParseLogLevel(conf.LogLevel))

	// init server
	srv := server.New(
		log,
		server.Config{
			GRPCPort: conf.GrpcPort,
			GRPCHost: conf.GrpcHost,
			HTTPPort: conf.HttpPort,
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
	err = shutdown.Wait(ctx, cancel, srv.Shutdown)
	if err != nil {
		log.Fatal("Error while shutting down server", err)
	}
}
