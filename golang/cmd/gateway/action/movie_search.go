package action

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg/grpc"
	pb "github.com/hidayatullahap/go-monorepo-example/pkg/proto/movies"
)

func (a *GatewayAction) MovieSearch(ctx context.Context, request entity.MovieSearchRequest) (entity.MovieList, error) {
	var res entity.MovieList
	conn, err := grpc.Dial(a.app.Config.Services.MovieHost)
	if err != nil {
		return res, err
	}

	defer conn.Close()

	pbReq := &pb.SearchRequest{
		Query: request.Search,
		Page:  request.Page,
	}

	pbMovies, err := pb.NewMoviesClient(conn).SearchMovie(ctx, pbReq)
	if err != nil {
		return res, err
	}

	res.Result = pbMovies.Result
	res.Page = pbMovies.Page
	res.TotalResults = pbMovies.TotalResult
	res.Movies = []entity.Movie{}

	for _, m := range pbMovies.Movies {
		res.Movies = append(res.Movies, entity.Movie{
			Title:  m.Title,
			Year:   m.Year,
			ImdbID: m.ImdbId,
			Type:   m.Type,
			Poster: m.Poster,
		})
	}

	return res, nil
}
