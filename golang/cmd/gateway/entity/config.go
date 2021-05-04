package entity

import (
	"os"
)

type Config struct {
	Services       Services
	HttpServerPort string
}

type Services struct {
	MovieHost string
	UserHost  string
}

func NewConfig() Config {
	var c Config

	c.Services = Services{
		MovieHost: os.Getenv("GRPC_HOST_MOVIE"),
		UserHost:  os.Getenv("GRPC_HOST_USER"),
	}

	c.HttpServerPort = os.Getenv("HTTP_PORT")

	return c
}
