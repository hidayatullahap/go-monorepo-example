package action

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/movie_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/cmd/movie_service/repo"
)

type IMovieAction interface {
	SearchMovie(ctx context.Context, search entity.SearchRequest) (entity.SearchResponse, error)
	DetailMovie(ctx context.Context, imdbID string) (entity.DetailResponse, error)
	Watchlist(ctx context.Context, req entity.WatchlistRequest) error
	GetWatchlist(ctx context.Context, userID string) ([]entity.Watchlist, error)
}

type MovieAction struct {
	movieRepo repo.IMovieRepo
}

func (a *MovieAction) SearchMovie(ctx context.Context, search entity.SearchRequest) (entity.SearchResponse, error) {
	return a.movieRepo.SearchMovie(ctx, search)
}

func (a *MovieAction) DetailMovie(ctx context.Context, imdbID string) (entity.DetailResponse, error) {
	return a.movieRepo.DetailMovie(ctx, imdbID)
}

func (a *MovieAction) Watchlist(ctx context.Context, req entity.WatchlistRequest) error {
	if req.Fav == false {
		errDelete := a.movieRepo.RemoveWatchlistMovie(ctx, req)
		return errDelete
	}

	err := a.movieRepo.AddWatchlistMovie(ctx, req)
	return err
}

func (a *MovieAction) GetWatchlist(ctx context.Context, userID string) ([]entity.Watchlist, error) {
	list, err := a.movieRepo.GetWatchlist(ctx, userID)
	return list, err
}

func NewMovieAction(app *entity.App) IMovieAction {
	return &MovieAction{
		movieRepo: repo.NewMovieRepo(app),
	}
}
