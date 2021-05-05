package action

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg/grpc"
	pb "github.com/hidayatullahap/go-monorepo-example/pkg/proto/users"
)

type IGatewayAction interface {
	Register(ctx context.Context, request entity.RegisterRequest) error
}

type GatewayAction struct {
	app *entity.App
}

func (a *GatewayAction) Register(ctx context.Context, request entity.RegisterRequest) error {
	conn, err := grpc.Dial("localhost:3001")
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

func NewGatewayAction(app *entity.App) *GatewayAction {
	return &GatewayAction{app}
}
