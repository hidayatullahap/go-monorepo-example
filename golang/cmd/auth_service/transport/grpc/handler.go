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
	userAction action.IAuthAction
}

func (h *Handler) Login(ctx context.Context, user *auth.User) (*auth.Token, error) {

	return &pb.Token{Token: "asd"}, nil
}

func NewGrpcHandler(app *entity.App) *Handler {
	return &Handler{
		app:        app,
		userAction: action.NewAuthAction(app),
	}
}
