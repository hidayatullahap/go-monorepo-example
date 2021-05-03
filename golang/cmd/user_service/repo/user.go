package repo

import (
	"context"
	"log"

	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepo interface {
	CreateUser(context.Context, entity.User) error
}

type UserRepo struct {
	db *mongo.Database
}

func (r *UserRepo) CreateUser(ctx context.Context, user entity.User) error {
	log.Print("didalam create user repo")

	return nil
}

func NewUserRepo(app *entity.App) IUserRepo {
	return &UserRepo{db: app.MongoDbClient}
}
