package mongo

type MongoConfig struct {
	Host     string `env:"DB_HOST,required"`
	Port     string `env:"DB_PORT,required"`
	Name     string `env:"DB_NAME,required"`
	Auth     string `env:"DB_AUTH"`
	Username string `env:"DB_USERNAME"`
	Password string `env:"DB_PASSWORD"`
	Pool     string `env:"DB_POOL"`
}
