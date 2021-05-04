package action

import (
	"context"
	"strings"

	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (a *UserAction) FindUser(ctx context.Context, username string) (entity.User, error) {
	var user entity.User
	username = strings.ToLower(username)

	filterUsername := bson.M{"username": username}
	user, err := a.userRepo.FindUser(ctx, filterUsername)
	if err != nil {
		if err == errors.ErrNotFound {
			err = errors.NotFound("user not found")
		}

		return user, err
	}

	return user, nil
}
