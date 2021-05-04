package action

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/movie_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/cmd/movie_service/repo"
)

type IMovieAction interface {
	SearchMovie(ctx context.Context, search entity.SearchRequest) (entity.SearchResponse, error)
	DetailMovie(ctx context.Context, omdbID string) (entity.DetailResponse, error)
}

type MovieAction struct {
	authRepo repo.IMovieRepo
}

func (a *MovieAction) SearchMovie(ctx context.Context, search entity.SearchRequest) (entity.SearchResponse, error) {
	return a.authRepo.SearchMovie(ctx, search)
}

func (a *MovieAction) DetailMovie(ctx context.Context, omdbID string) (entity.DetailResponse, error) {
	return a.authRepo.DetailMovie(ctx, omdbID)
}

func NewMovieAction(app *entity.App) IMovieAction {
	return &MovieAction{
		authRepo: repo.NewMovieRepo(app),
	}
}
