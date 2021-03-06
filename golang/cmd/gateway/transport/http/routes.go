package http

import (
	"net/http"

	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/entity"
	"github.com/labstack/echo/v4"
)

func setupRoutes(e *echo.Echo, app *entity.App) {
	e.GET("/", noContent)

	h := NewHandler(app)
	m := NewMiddlewareAuth(app.Config)

	v1 := e.Group("/api/v1")
	v1.POST("/users/register", h.Register)
	v1.POST("/users/login", h.Login)

	v1.Use(m.AuthCheckToken)
	v1.GET("/watchlist", h.GetWatchlist)
	v1.PUT("/watchlist/:imdb_id", h.AddToWatchlist)
	v1.DELETE("/watchlist/:imdb_id", h.RemoveFromWatchlist)

	v1.GET("/movies", h.MovieSearch)
}

func noContent(e echo.Context) error {
	return e.NoContent(http.StatusOK)
}
