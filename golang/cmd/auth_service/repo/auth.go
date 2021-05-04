package repo

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/auth_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg/grpc"
	pb "github.com/hidayatullahap/go-monorepo-example/pkg/proto/users"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAuthRepo interface {
	FindUser(ctx context.Context, username string) (entity.User, error)
}

type AuthRepo struct {
	db    *mongo.Database
	hosts entity.Services
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
