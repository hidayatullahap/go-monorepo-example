package repo

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepo interface {
	CreateUser(context.Context, entity.User) error
}

type UserRepo struct {
	db *mongo.Database
}

func (r *UserRepo) CreateUser(ctx context.Context, user entity.User) error {
	user.ID = pkg.NewULID()

	_, err := r.db.Collection(CollectionUsers).InsertOne(ctx, user)
	return err
}

func NewUserRepo(app *entity.App) IUserRepo {
	return &UserRepo{db: app.MongoDbClient}
}
