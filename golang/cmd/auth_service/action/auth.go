package action

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/auth_service/repo"
	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/entity"
)

type IAuthAction interface {
	Login(ctx context.Context, user entity.User) (string, error)
}

type AuthAction struct {
	authRepo repo.IAuthRepo
}

func NewAuthAction(app *entity.App) IAuthAction {
	return &AuthAction{
		authRepo: repo.NewAuthRepo(app),
	}
}
