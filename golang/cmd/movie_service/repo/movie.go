package repo

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/hidayatullahap/go-monorepo-example/cmd/movie_service/entity"
	"github.com/parnurzeal/gorequest"
	"go.mongodb.org/mongo-driver/mongo"
)

type IMovieRepo interface {
	SearchMovie(ctx context.Context, search entity.SearchRequest) (entity.SearchResponse, error)
}

type MovieRepo struct {
	db      *mongo.Database
	service entity.Services
}

func (r *MovieRepo) SearchMovie(ctx context.Context, search entity.SearchRequest) (entity.SearchResponse, error) {
	_, body, _ := gorequest.New().
		Get(r.service.OmdbHost).
		Query("apikey=" + r.service.OmdbApiKey).
		Query("s=" + url.QueryEscape(search.Query)).
		Query("page=" + strconv.Itoa(int(search.Page))).
		End()

	var res entity.SearchResponse
	err := json.Unmarshal([]byte(body), &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func NewMovieRepo(app *entity.App) IMovieRepo {
	return &MovieRepo{
		db:      app.MongoDbClient,
		service: app.Config.Services,
	}
}
