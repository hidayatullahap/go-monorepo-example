package grpc

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/movie_service/action"
	"github.com/hidayatullahap/go-monorepo-example/cmd/movie_service/builder"
	"github.com/hidayatullahap/go-monorepo-example/cmd/movie_service/entity"
	pb "github.com/hidayatullahap/go-monorepo-example/pkg/proto/movies"
)

type Handler struct {
	app         *entity.App
	movieAction action.IMovieAction
}

func (h *Handler) GetWatchlist(ctx context.Context, request *pb.GetWatchlistRequest) (*pb.WatchlistResponse, error) {
	list, err := h.movieAction.GetWatchlist(ctx, request.UserId)

	var res pb.WatchlistResponse
	var movies []*pb.WatchlistMovie
	for _, movie := range list {
		movies = append(movies, &pb.WatchlistMovie{
			Id:         movie.ID,
			UserId:     movie.UserID,
			OmdbId:     movie.OmdbID,
			MovieTitle: movie.MovieTitle,
		})
	}

	res.Movies = movies

	return &res, err
}

func (h *Handler) Watchlist(ctx context.Context, request *pb.WatchlistRequest) (*pb.NoResponse, error) {
	err := h.movieAction.Watchlist(ctx, entity.WatchlistRequest{
		UserID: request.UserId,
		OmdbID: request.OmdbId,
		Fav:    request.Fav,
	})

	return &pb.NoResponse{}, err
}

func (h *Handler) DetailMovie(ctx context.Context, request *pb.DetailRequest) (*pb.DetailResponse, error) {
	res, err := h.movieAction.DetailMovie(ctx, request.OmdbId)
	if err != nil {
		return nil, err
	}

	pbRes := builder.BuildDetailMovieProto(res)
	return pbRes, err
}

func (h *Handler) SearchMovie(ctx context.Context, request *pb.SearchRequest) (*pb.SearchResponse, error) {
	req := entity.SearchRequest{
		Query: request.Query,
		Page:  request.Page,
	}

	res, err := h.movieAction.SearchMovie(ctx, req)
	if err != nil {
		return nil, err
	}

	pbRes := builder.BuildSearchProto(res)
	pbRes.Page = request.Page

	return pbRes, err
}

func NewGrpcHandler(app *entity.App) *Handler {
	return &Handler{
		app:         app,
		movieAction: action.NewMovieAction(app),
	}
}
