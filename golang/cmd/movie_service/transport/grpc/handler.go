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
