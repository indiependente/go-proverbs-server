package main

import "github.com/caarlos0/env/v7"

type config struct {
	GrpcPort string `env:"GRPC_PORT" envDefault:"9090"`
	GrpcHost string `env:"GRPC_HOST" envDefault:"localhost"`
	HttpPort string `env:"HTTP_PORT" envDefault:"8080"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"INFO"`
}

func parseConfig() (*config, error) {
	var c config

	return &c, env.Parse(&c)
}
