package grpc

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/entity"
	pb "github.com/hidayatullahap/go-monorepo-example/pkg/proto/user"
)

type Handler struct {
	app *entity.App
}

func (h Handler) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Msg: request.Msg}, nil
}

func NewGrpcHandler(app *entity.App) *Handler {
	return &Handler{
		app: app,
	}
}
