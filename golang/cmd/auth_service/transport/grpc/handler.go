package grpc

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/auth_service/action"
	"github.com/hidayatullahap/go-monorepo-example/cmd/auth_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg/proto/auth"
	pb "github.com/hidayatullahap/go-monorepo-example/pkg/proto/auth"
)

type Handler struct {
	app        *entity.App
	authAction action.IAuthAction
}

func (h *Handler) Auth(ctx context.Context, token *pb.Token) (*pb.User, error) {
	user, err := h.authAction.Auth(ctx, token.Token)
	if err != nil {
		return nil, err
	}

	res := &pb.User{
		Id:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
	}

	return res, nil
}

func (h *Handler) Login(ctx context.Context, user *auth.User) (*auth.Token, error) {
	userRequest := entity.User{
		Username: user.Username,
		Password: user.Password,
	}

	token, err := h.authAction.Login(ctx, userRequest)
	if err != nil {
		return nil, err
	}

	return &pb.Token{Token: token}, nil
}

func NewGrpcHandler(app *entity.App) *Handler {
	return &Handler{
		app:        app,
		authAction: action.NewAuthAction(app),
	}
}
