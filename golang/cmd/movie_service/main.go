package main

import (
	"context"
	"log"
	"sync"

	"github.com/hidayatullahap/go-monorepo-example/cmd/movie_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/cmd/movie_service/transport"
	"github.com/hidayatullahap/go-monorepo-example/pkg/db/mongo"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := entity.NewApp()
	mongoClient, err := mongo.NewMongo(app.Config.Mongo).Connect(ctx)
	if err != nil {
		panic(err)
	}

	app.MongoDbClient = mongoClient

	t := transport.NewTransport(app)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	t.GrpcServer.Start()

	wg.Wait()
}
