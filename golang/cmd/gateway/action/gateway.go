package action

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/entity"
	m "github.com/hidayatullahap/go-monorepo-example/cmd/movie_service/entity"
)

type IGatewayAction interface {
	Register(ctx context.Context, request entity.RegisterRequest) error
	Login(ctx context.Context, request entity.LoginRequest) (string, error)
	MovieSearch(ctx context.Context, request entity.MovieSearchRequest) (entity.MovieList, error)
	MovieDetail(ctx context.Context, request entity.MovieSearchRequest) (entity.MovieDetail, error)
	Watchlist(ctx context.Context, request m.WatchlistRequest) error
	UserWatchlist(ctx context.Context, userID string) ([]m.Watchlist, error)
}

type GatewayAction struct {
	app *entity.App
}

func NewGatewayAction(app *entity.App) *GatewayAction {
	return &GatewayAction{app}
}
