package action

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg/grpc"
	pb "github.com/hidayatullahap/go-monorepo-example/pkg/proto/users"
)

func (a *GatewayAction) Register(ctx context.Context, request entity.RegisterRequest) error {
	conn, err := grpc.Dial(a.app.Config.Services.UserHost)
	if err != nil {
		return err
	}

	defer conn.Close()

	pbReq := &pb.User{
		Username: request.Username,
		FullName: request.FullName,
		Password: request.Password,
	}

	_, err = pb.NewUsersClient(conn).CreateUser(ctx, pbReq)
	if err != nil {
		return err
	}

	return nil
}
