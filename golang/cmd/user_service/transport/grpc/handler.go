package grpc

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/entity"
	pb "github.com/hidayatullahap/go-monorepo-example/pkg/proto/users"
)

type Handler struct {
	app *entity.App
}

func (h Handler) CreateUser(ctx context.Context, user *pb.User) (*pb.NoResponse, error) {
	return &pb.NoResponse{}, nil
}

func (h Handler) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Msg: request.Msg}, nil
}

func NewGrpcHandler(app *entity.App) *Handler {
	return &Handler{
		app: app,
	}
}
