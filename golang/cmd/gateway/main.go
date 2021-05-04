package main

import (
	"log"
	"sync"

	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/entity"
	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/transport"
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

	t.HttpServer.Start()

	wg.Wait()
}
