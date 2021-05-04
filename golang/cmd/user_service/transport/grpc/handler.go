package grpc

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/action"
	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/builder"
	"github.com/hidayatullahap/go-monorepo-example/cmd/user_service/entity"
	pb "github.com/hidayatullahap/go-monorepo-example/pkg/proto/users"
)

type Handler struct {
	app        *entity.App
	userAction action.IUserAction
}

func (h *Handler) FindUser(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
	user, err := h.userAction.FindUser(ctx, entity.User{
		ID:       req.UserId,
		Username: req.Username,
	})

	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		Password: user.Password,
	}, nil
}

func (h *Handler) CreateUser(ctx context.Context, pbUser *pb.User) (*pb.NoResponse, error) {
	user := builder.BuildUserFromProto(pbUser)
	err := h.userAction.CreateUser(ctx, user)

	return &pb.NoResponse{}, err
}

func (h *Handler) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Msg: request.Msg}, nil
}

func NewGrpcHandler(app *entity.App) *Handler {
	return &Handler{
		app:        app,
		userAction: action.NewUserAction(app),
	}
}
