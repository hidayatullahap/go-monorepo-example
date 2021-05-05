package http

import (
	netHttp "net/http"

	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/action"
	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg/http"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	app    *entity.App
	action action.IGatewayAction
}

func (h *Handler) Register(c echo.Context) error {
	var req entity.RegisterRequest
	err := http.BindAndValidate(c, &req)
	if err != nil {
		return err
	}

	err = h.action.Register(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.NoContent(netHttp.StatusCreated)
}

func NewHandler(app *entity.App) *Handler {
	return &Handler{
		app:    app,
		action: action.NewGatewayAction(app),
	}
}
