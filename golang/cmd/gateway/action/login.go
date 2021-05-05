package action

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg/grpc"
	pb "github.com/hidayatullahap/go-monorepo-example/pkg/proto/auth"
)

func (a *GatewayAction) Login(ctx context.Context, request entity.LoginRequest) (string, error) {
	var token string

	conn, err := grpc.Dial(a.app.Config.Services.AuthHost)
	if err != nil {
		return token, err
	}

	defer conn.Close()

	pbReq := &pb.User{
		Username: request.Username,
		Password: request.Password,
	}

	pbRes, err := pb.NewAuthClient(conn).Login(ctx, pbReq)
	if err != nil {
		return token, err
	}

	token = pbRes.Token
	return token, nil
}
