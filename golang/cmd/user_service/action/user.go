package action

import (
	"context"
	"strings"

	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/repo"
	"github.com/hidayatullahap/go-monorepo-example/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

type IUserAction interface {
	CreateUser(context.Context, entity.User) error
}

type UserAction struct {
	userRepo repo.IUserRepo
}

func (a *UserAction) CreateUser(ctx context.Context, user entity.User) error {
	user.Username = strings.ToLower(user.Username)

	filterUsername := bson.M{"username": user.Username}
	dUser, err := a.userRepo.FindUser(ctx, filterUsername)
	if err != nil && err != errors.ErrNotFound {
		return err
	}

	if dUser.Username == user.Username {
		return errors.InvalidArgument("username already used, please choose another one")
	}

	err = a.userRepo.CreateUser(ctx, user)
	return err
}

func NewUserAction(app *entity.App) IUserAction {
	return &UserAction{
		userRepo: repo.NewUserRepo(app),
	}
}
