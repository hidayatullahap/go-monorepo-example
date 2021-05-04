package action

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/movie_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/cmd/movie_service/repo"
)

type IMovieAction interface {
	SearchMovie(ctx context.Context, search entity.SearchRequest) (entity.SearchResponse, error)
}

type MovieAction struct {
	authRepo repo.IMovieRepo
}

func (a *MovieAction) SearchMovie(ctx context.Context, search entity.SearchRequest) (entity.SearchResponse, error) {
	return a.authRepo.SearchMovie(ctx, search)
}

func NewMovieAction(app *entity.App) IMovieAction {
	return &MovieAction{
		authRepo: repo.NewMovieRepo(app),
	}
}
