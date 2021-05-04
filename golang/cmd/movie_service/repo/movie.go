package repo

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hidayatullahap/go-monorepo-example/cmd/movie_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg"
	m "github.com/hidayatullahap/go-monorepo-example/pkg/db/mongo"
	"github.com/hidayatullahap/go-monorepo-example/pkg/errors"
	"github.com/parnurzeal/gorequest"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMovieRepo interface {
	SearchMovie(ctx context.Context, search entity.SearchRequest) (entity.SearchResponse, error)
	DetailMovie(ctx context.Context, omdbID string) (entity.DetailResponse, error)
	AddWatchlistMovie(ctx context.Context, req entity.WatchlistRequest) error
	RemoveWatchlistMovie(ctx context.Context, req entity.WatchlistRequest) error
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

func (r *MovieRepo) DetailMovie(ctx context.Context, omdbID string) (entity.DetailResponse, error) {
	_, body, _ := gorequest.New().
		Get(r.service.OmdbHost).
		Query("apikey=" + r.service.OmdbApiKey).
		Query("i=" + omdbID).
		End()

	var res entity.DetailResponse
	err := json.Unmarshal([]byte(body), &res)
	if err != nil {
		return res, err
	}

	if res.Error != "" {
		return res, errors.InvalidArgument(res.Error)
	}

	return res, nil
}

func (r *MovieRepo) AddWatchlistMovie(ctx context.Context, req entity.WatchlistRequest) error {
	movie, err := r.DetailMovie(ctx, req.OmdbID)
	if err != nil {
		return err
	}

	wl := entity.Watchlist{
		UserID:     req.UserID,
		OmdbID:     req.OmdbID,
		MovieTitle: movie.Title,
	}

	opt := &options.UpdateOptions{
		Upsert: aws.Bool(true),
	}

	data := bson.M{"$set": &wl, "$setOnInsert": bson.M{"_id": pkg.NewULID()}}
	_, err = r.db.Collection(m.CollectionWatchlist).UpdateOne(ctx, bson.M{"user_id": req.UserID}, &data, opt)
	return err
}

func (r *MovieRepo) RemoveWatchlistMovie(ctx context.Context, req entity.WatchlistRequest) error {
	filter := bson.M{"omdb_id": req.OmdbID, "user_id": req.UserID}
	_, err := r.db.Collection(m.CollectionWatchlist).DeleteOne(ctx, filter)
	return err
}

func NewMovieRepo(app *entity.App) IMovieRepo {
	return &MovieRepo{
		db:      app.MongoDbClient,
		service: app.Config.Services,
	}
}
