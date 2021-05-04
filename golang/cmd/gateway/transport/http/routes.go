package http

import (
	"net/http"

	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/entity"
	"github.com/labstack/echo/v4"
)

func setupRoutes(e *echo.Echo, app *entity.App) {
	e.GET("/", noContent)
}

func noContent(e echo.Context) error {
	return e.NoContent(http.StatusOK)
}
