package action

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/repo"
)

type IUserAction interface {
	CreateUser(context.Context, entity.User) error
}

type UserAction struct {
	userRepo repo.IUserRepo
}

// TODO: search existing username
func (a *UserAction) CreateUser(ctx context.Context, user entity.User) error {
	err := a.userRepo.CreateUser(ctx, user)
	return err
}

func NewUserAction(app *entity.App) IUserAction {
	return &UserAction{
		userRepo: repo.NewUserRepo(app),
	}
}
