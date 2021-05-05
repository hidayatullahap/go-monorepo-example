package action

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg/grpc"
	pb "github.com/hidayatullahap/go-monorepo-example/pkg/proto/movies"
)

func (a *GatewayAction) MovieDetail(ctx context.Context, request entity.MovieSearchRequest) (entity.MovieDetail, error) {
	var res entity.MovieDetail
	conn, err := grpc.Dial(a.app.Config.Services.MovieHost)
	if err != nil {
		return res, err
	}

	defer conn.Close()

	pbReq := &pb.DetailRequest{ImdbId: request.ImdbID}
	m, err := pb.NewMoviesClient(conn).DetailMovie(ctx, pbReq)
	if err != nil {
		return res, err
	}

	res = entity.MovieDetail{
		Title:      m.Title,
		Year:       m.Year,
		Rated:      m.Rated,
		Released:   m.Released,
		Runtime:    m.Runtime,
		Genre:      m.Genre,
		Director:   m.Director,
		Writer:     m.Writer,
		Actors:     m.Actors,
		Plot:       m.Plot,
		Language:   m.Language,
		Country:    m.Country,
		Awards:     m.Awards,
		Poster:     m.Poster,
		Ratings:    []entity.Rating{},
		Metascore:  m.Metascore,
		ImdbRating: m.ImdbRating,
		ImdbVotes:  m.ImdbVotes,
		ImdbID:     m.ImdbId,
		Type:       m.Type,
		DVD:        m.DVD,
		BoxOffice:  m.BoxOffice,
		Production: m.Production,
		Website:    m.Website,
	}

	for _, rating := range m.Ratings {
		res.Ratings = append(res.Ratings, entity.Rating{
			Source: rating.Source,
			Value:  rating.Value,
		})
	}

	return res, nil
}
