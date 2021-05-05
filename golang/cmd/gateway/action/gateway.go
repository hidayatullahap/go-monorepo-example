package action

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/entity"
)

type IGatewayAction interface {
	Register(ctx context.Context, request entity.RegisterRequest) error
}

type GatewayAction struct {
	app *entity.App
}

func NewGatewayAction(app *entity.App) *GatewayAction {
	return &GatewayAction{app}
}
