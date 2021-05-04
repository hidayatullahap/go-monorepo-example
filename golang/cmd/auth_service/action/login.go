package action

import (
	"context"
	"strings"

	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg"
	"github.com/hidayatullahap/go-monorepo-example/pkg/errors"
)

func (a *AuthAction) Login(ctx context.Context, user entity.User) (string, error) {
	var token string
	user.Username = strings.ToLower(user.Username)

	dUser, err := a.authRepo.FindUser(ctx, user.Username)
	if err != nil {
		if err == errors.ErrNotFound {
			err = errors.InvalidArgument("username or password not match")
		}

		return token, err
	}

	match := pkg.ComparePasswords(dUser.Password, []byte(user.Password))
	if !match {
		return token, errors.InvalidArgument("username or password not match")
	}

	token, err = pkg.GenerateToken(user.Username)
	if err != nil {
		return token, err
	}

	return token, nil
}
