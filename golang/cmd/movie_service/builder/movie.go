package builder

import (
	"strconv"

	"github.com/hidayatullahap/go-monorepo-example/cmd/movie_service/entity"
	pb "github.com/hidayatullahap/go-monorepo-example/pkg/proto/movies"
)

func BuildSearchProto(search entity.SearchResponse) *pb.SearchResponse {
	var res pb.SearchResponse
	var movies []*pb.Movie

	for _, movie := range search.Movies {
		movies = append(movies, &pb.Movie{
			Title:  movie.Title,
			Year:   movie.Year,
			ImdbId: movie.ImdbID,
			Type:   movie.Type,
			Poster: movie.Poster,
		})
	}

	res.Movies = movies

	total, _ := strconv.Atoi(search.TotalResults)
	res.TotalResult = int64(total)
	res.Result = int64(len(movies))

	return &res
}
