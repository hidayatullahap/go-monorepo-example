package action

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (a *UserAction) FindUser(ctx context.Context, req entity.User) (entity.User, error) {
	var user entity.User

	filter := a.getFilter(req)
	user, err := a.userRepo.FindUser(ctx, filter)
	if err != nil {
		if err == errors.ErrNotFound {
			err = errors.ErrUserNotFound
		}

		return user, err
	}

	return user, nil
}

func (a *UserAction) getFilter(req entity.User) bson.M {
	filter := bson.M{}

	if req.ID != "" {
		filter["_id"] = req.ID

		return filter
	}

	if req.Username != "" {
		filter["username"] = req.Username

		return filter
	}

	return filter
}
