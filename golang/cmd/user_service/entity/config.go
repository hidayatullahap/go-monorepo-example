package entity

import "os"

type Mongo struct {
	Host     string `env:"DB_HOST,required"`
	Port     string `env:"DB_PORT,required"`
	Name     string `env:"DB_NAME,required"`
	Auth     string `env:"DB_AUTH"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Pool     string `env:"DB_POOL"`
}

type Config struct {
	Mongo Mongo
}

func NewConfig() Config {
	var c Config

	c.Mongo = Mongo{
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
