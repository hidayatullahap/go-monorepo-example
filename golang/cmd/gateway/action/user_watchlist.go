package action

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/movie_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg/grpc"
	pb "github.com/hidayatullahap/go-monorepo-example/pkg/proto/movies"
)

func (a *GatewayAction) UserWatchlist(ctx context.Context, userID string) ([]entity.Watchlist, error) {
	list := []entity.Watchlist{}

	conn, err := grpc.Dial(a.app.Config.Services.MovieHost)
	if err != nil {
		return list, err
	}

	defer conn.Close()

	pbReq := &pb.GetWatchlistRequest{UserId: userID}
	pbWatchlist, err := pb.NewMoviesClient(conn).GetWatchlist(ctx, pbReq)
	if err != nil {
		return list, err
	}

	for _, movie := range pbWatchlist.Movies {
		list = append(list, entity.Watchlist{
			ID:         movie.Id,
			UserID:     movie.UserId,
			ImdbID:     movie.ImdbId,
			MovieTitle: movie.MovieTitle,
		})
	}

	return list, nil
}
