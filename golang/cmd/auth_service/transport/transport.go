package transport

import (
	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/transport/grpc"
)

type Transport struct {
	GrpcServer *grpc.Server
}

func NewTransport(app *entity.App) *Transport {
	return &Transport{
		grpc.NewGrpcServer(app),
	}
}
