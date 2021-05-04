package action

import (
	"context"
	"strings"

	"github.com/hidayatullahap/go-monorepo-example/cmd/auth_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg"
	"github.com/hidayatullahap/go-monorepo-example/pkg/errors"
	"github.com/hidayatullahap/go-monorepo-example/pkg/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *AuthAction) Login(ctx context.Context, user entity.User) (string, error) {
	var token string
	user.Username = strings.ToLower(user.Username)

	dUser, err := a.authRepo.FindUser(ctx, user)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			if st.Code() == codes.NotFound {
				err = errors.InvalidArgument("username or password not match")
			}
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

	upsertToken := entity.Token{
		UserID:   dUser.ID,
		Username: user.Username,
		Token:    token,
	}

	err = a.authRepo.UpdateToken(ctx, upsertToken)
	if err != nil {
		return token, err
	}

	return token, nil
}
