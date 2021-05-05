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
	AuthHost  string
}

func NewConfig() Config {
	var c Config

	c.Services = Services{
		MovieHost: os.Getenv("HOST_GRPC_MOVIE_SERVICE"),
		UserHost:  os.Getenv("HOST_GRPC_USER_SERVICE"),
		AuthHost:  os.Getenv("HOST_GRPC_AUTH_SERVICE"),
	}

	c.HttpServerPort = os.Getenv("HTTP_PORT")

	return c
}
