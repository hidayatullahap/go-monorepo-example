package action

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/auth_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/cmd/auth_service/repo"
	"github.com/hidayatullahap/go-monorepo-example/pkg/errors"
)

type IAuthAction interface {
	Login(ctx context.Context, user entity.User) (string, error)
	Auth(ctx context.Context, token string) (entity.User, error)
}

type AuthAction struct {
	authRepo repo.IAuthRepo
}

func (a *AuthAction) Auth(ctx context.Context, token string) (entity.User, error) {
	userToken, err := a.authRepo.FindUserToken(ctx, token)
	if err != nil {
		if err == errors.ErrNotFound {
			err = errors.ErrUnauthorized
		}

		return entity.User{}, err
	}

	req := entity.User{
		Username: userToken.Username,
		ID:       userToken.UserID,
	}

	res, err := a.authRepo.FindUser(ctx, req)
	if err != nil {
		return res, err
	}

	return res, nil
}

func NewAuthAction(app *entity.App) IAuthAction {
	return &AuthAction{
		authRepo: repo.NewAuthRepo(app),
	}
}
