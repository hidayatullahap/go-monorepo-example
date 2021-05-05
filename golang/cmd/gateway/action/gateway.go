package action

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/entity"
)

type IGatewayAction interface {
	Register(ctx context.Context, request entity.RegisterRequest) error
	Login(ctx context.Context, request entity.LoginRequest) (string, error)
	MovieSearch(ctx context.Context, request entity.MovieSearchRequest) (entity.MovieList, error)
}

type GatewayAction struct {
	app *entity.App
}

func NewGatewayAction(app *entity.App) *GatewayAction {
	return &GatewayAction{app}
}
