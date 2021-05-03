package builder

import (
	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg/proto/users"
)

func BuildUserFromProto(user *users.User) entity.User {
	return entity.User{
		Username: user.Username,
		FullName: user.FullName,
		Password: user.Password,
	}
}
