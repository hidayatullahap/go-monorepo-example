package action

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/repo"
)

type IUserAction interface {
	CreateUser(context.Context, entity.User) error
	Login(ctx context.Context, user entity.User) (string, error)
}

type UserAction struct {
	userRepo repo.IUserRepo
}

func NewUserAction(app *entity.App) IUserAction {
	return &UserAction{
		userRepo: repo.NewUserRepo(app),
	}
}
