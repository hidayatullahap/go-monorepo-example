package repo

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg"
	m "github.com/hidayatullahap/go-monorepo-example/pkg/db/mongo"
	"github.com/hidayatullahap/go-monorepo-example/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepo interface {
	CreateUser(context.Context, entity.User) error
	FindUser(ctx context.Context, filter bson.M) (entity.User, error)
}

type UserRepo struct {
	db    *mongo.Database
	crypt *pkg.Crypt
}

func (r *UserRepo) CreateUser(ctx context.Context, user entity.User) error {
	var hashedPassword string
	hashedPassword, err := r.crypt.HashAndSalt([]byte(user.Password))
	if err != nil {
		return errors.InternalError(err.Error())
	}

	user.ID = pkg.NewULID()
	user.Password = hashedPassword

	_, err = r.db.Collection(m.CollectionUsers).InsertOne(ctx, user)
	return err
}

func (r *UserRepo) FindUser(ctx context.Context, filter bson.M) (entity.User, error) {
	var user entity.User
	err := r.db.Collection(m.CollectionUsers).FindOne(ctx, filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		err = errors.ErrNotFound
	}

	return user, err
}

func NewUserRepo(app *entity.App) IUserRepo {
	return &UserRepo{db: app.MongoDbClient}
}
