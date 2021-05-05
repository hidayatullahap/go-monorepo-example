package action

import (
	"context"

	"github.com/hidayatullahap/go-monorepo-example/cmd/movie_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg/grpc"
	pb "github.com/hidayatullahap/go-monorepo-example/pkg/proto/movies"
)

func (a *GatewayAction) Watchlist(ctx context.Context, request entity.WatchlistRequest) error {
	conn, err := grpc.Dial(a.app.Config.Services.MovieHost)
	if err != nil {
		return err
	}

	defer conn.Close()

	pbReq := &pb.WatchlistRequest{
		ImdbId: request.ImdbID,
		Fav:    request.Fav,
		UserId: request.UserID,
	}

	_, err = pb.NewMoviesClient(conn).Watchlist(ctx, pbReq)
	if err != nil {
		return err
	}

	return nil
}
