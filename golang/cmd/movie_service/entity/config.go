package entity

import (
	"os"

	"github.com/hidayatullahap/go-monorepo-example/pkg/db/mongo"
)

type Config struct {
	Mongo    mongo.MongoConfig
	Services Services
}

type Services struct {
	OmdbHost   string
	OmdbApiKey string
}

func NewConfig() Config {
	var c Config

	c.Mongo = mongo.MongoConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		Auth:     os.Getenv("DB_AUTH"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Pool:     os.Getenv("DB_POOL"),
	}

	c.Services = Services{
		OmdbHost:   os.Getenv("OMBD_HOST"),
		OmdbApiKey: os.Getenv("OMBD_API_KEY"),
	}

	return c
}
