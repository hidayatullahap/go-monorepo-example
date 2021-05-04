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

func BuildDetailMovieProto(movie entity.DetailResponse) *pb.DetailResponse {
	res := pb.DetailResponse{
		Year:       movie.Year,
		Rated:      movie.Rated,
		Released:   movie.Released,
		Runtime:    movie.Runtime,
		Genre:      movie.Genre,
		Director:   movie.Director,
		Writer:     movie.Writer,
		Actors:     movie.Actors,
		Plot:       movie.Plot,
		Language:   movie.Language,
		Country:    movie.Country,
		Awards:     movie.Awards,
		Poster:     movie.Poster,
		Metascore:  movie.Metascore,
		ImdbRating: movie.ImdbRating,
		ImdbVotes:  movie.ImdbVotes,
		ImdbId:     movie.ImdbID,
		Type:       movie.Type,
		DVD:        movie.DVD,
		BoxOffice:  movie.BoxOffice,
		Production: movie.Production,
		Website:    movie.Website,
		Response:   movie.Response,
		Title:      movie.Title,
	}

	var ratings []*pb.Rating
	for _, r := range movie.Ratings {
		ratings = append(ratings, &pb.Rating{
			Source: r.Source,
			Value:  r.Value,
		})
	}

	res.Ratings = ratings

	return &res
}
