package repo

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hidayatullahap/go-monorepo-example/cmd/auth_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg"
	m "github.com/hidayatullahap/go-monorepo-example/pkg/db/mongo"
	"github.com/hidayatullahap/go-monorepo-example/pkg/grpc"
	pb "github.com/hidayatullahap/go-monorepo-example/pkg/proto/users"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IAuthRepo interface {
	FindUser(ctx context.Context, username string) (entity.User, error)
	UpdateToken(ctx context.Context, token entity.Token) error
}

type AuthRepo struct {
	db    *mongo.Database
	hosts entity.Services
}

func (r *AuthRepo) UpdateToken(ctx context.Context, token entity.Token) error {
	opt := &options.UpdateOptions{
		Upsert: aws.Bool(true),
	}

	data := bson.M{"$set": &token, "$setOnInsert": bson.M{"_id": pkg.NewULID()}}
	_, err := r.db.Collection(m.CollectionUserToken).UpdateOne(ctx, bson.M{"user_id": token.UserID}, &data, opt)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthRepo) FindUser(ctx context.Context, username string) (entity.User, error) {
	var user entity.User

	conn, err := grpc.Dial(r.hosts.UserServiceHost)
	if err != nil {
		return user, err
	}

	defer conn.Close()

	res, err := pb.NewUsersClient(conn).FindUser(ctx, &pb.UserRequest{Username: username})
	if err != nil {
		return user, err
	}

	user.ID = res.Id
	user.Username = res.Username
	user.Password = res.Password
	user.FullName = res.FullName

	return user, nil
}

func NewAuthRepo(app *entity.App) IAuthRepo {
	return &AuthRepo{
		db:    app.MongoDbClient,
		hosts: app.Config.Services,
	}
}
