package entity

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Config        Config
	MongoDbClient *mongo.Database
}

func NewApp() *App {
	app := &App{Config: NewConfig()}
	return app
}
