package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"go.elastic.co/apm/module/apmmongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	config MongoConfig
}

func (m *Mongo) Connect(ctx context.Context) (*mongo.Database, error) {
	var urlConnection string

	urlConnection = m.devConnectionScript()
	mode := "dev"
	envMode := os.Getenv("APP_ENV")
	if envMode != "" {
		mode = envMode
	}

	if mode == "production" || mode == "prod" {
		urlConnection = m.prodConnectionScript()
	}

	// default pool number
	maxPool := 100
	maxPoolI, _ := strconv.Atoi(m.config.Pool)
	if maxPoolI > 0 {
		maxPool = maxPoolI
	}

	opt := options.Client().SetAppName("movie-app").
		SetMonitor(apmmongo.CommandMonitor()).
		SetMaxPoolSize(uint64(maxPool))

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(urlConnection), opt)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error connecting to MongoDB: %s", err.Error()))
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Printf("Connected to MongoDB! env: %v pool size: %v", mode, maxPool)
	db := client.Database(m.config.Name)
	return db, nil
}

func (m *Mongo) devConnectionScript() string {
	hostName := fmt.Sprintf("%s:%s", m.config.Host, m.config.Port)

	var usernameAndPassword string
	if m.config.Username != "" {
		usernameAndPassword = fmt.Sprintf("%s:%s@", m.config.Username, m.config.Password)
	}

	// auth source database for mongodb connection, if empty it would be default as `admin`
	var authSource string
	if m.config.Auth != "" {
		authSource = "/?authSource=" + m.config.Auth
	}

	log.Println("Connecting to MongoDB Server Dev " + hostName + " database: " + m.config.Name)
	return fmt.Sprintf("mongodb://%s%s%s", usernameAndPassword, hostName, authSource)
}

func (m *Mongo) prodConnectionScript() string {
	hostName := fmt.Sprintf("%s/%s", m.config.Host, m.config.Name)
	log.Println("Connecting to MongoDB Server Prod " + hostName + " ...")

	var usernameAndPassword string
	if m.config.Username != "" {
		usernameAndPassword = fmt.Sprintf("%s:%s@", m.config.Username, m.config.Password)
	}

	return fmt.Sprintf("mongodb+srv://%s%s", usernameAndPassword, hostName)
}

func NewMongo(config MongoConfig) *Mongo {
	return &Mongo{
		config: config,
	}
}
