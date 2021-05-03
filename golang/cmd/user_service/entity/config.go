package entity

import (
	"github.com/hidayatullahap/go-monorepo-example/pkg"
)

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
	var (
		c     Config
		mongo Mongo
	)

	var configs []interface{}
	configs = append(configs, &mongo)
	for _, i := range configs {
		err := pkg.MarshalEnv(i)
		if err != nil {
			panic(err)
		}
	}

	c.Mongo = mongo

	return c
}
