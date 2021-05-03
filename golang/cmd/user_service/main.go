package main

import (
	"log"
	"sync"

	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/transport"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := entity.NewApp()

	t := transport.NewTransport(app)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	t.GrpcServer.Start()

	wg.Wait()
}
