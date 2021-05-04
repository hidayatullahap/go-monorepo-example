package entity

import (
	"os"

	"github.com/hidayatullahap/go-monorepo-example/pkg/db/mongo"
)

type Config struct {
	Mongo mongo.MongoConfig
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

	return c
}
