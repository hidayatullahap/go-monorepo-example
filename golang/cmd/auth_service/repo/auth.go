package repo

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/auth_service/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAuthRepo interface {
	FindUser(ctx context.Context, username string) (entity.User, error)
}

type AuthRepo struct {
	db *mongo.Database
}

func (r *AuthRepo) FindUser(ctx context.Context, username string) (entity.User, error) {
	// TODO: get from proto
	return entity.User{}, nil
}

func NewAuthRepo(app *entity.App) IAuthRepo {
	return &AuthRepo{db: app.MongoDbClient}
}
