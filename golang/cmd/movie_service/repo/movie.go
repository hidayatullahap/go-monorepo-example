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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/resty.v1"
)

type IMovieRepo interface {
	SearchMovie(ctx context.Context, search entity.SearchRequest) (entity.SearchResponse, error)
	DetailMovie(ctx context.Context, imdbID string) (entity.DetailResponse, error)
	AddWatchlistMovie(ctx context.Context, req entity.WatchlistRequest) error
	RemoveWatchlistMovie(ctx context.Context, req entity.WatchlistRequest) error
	GetWatchlist(ctx context.Context, userID string) ([]entity.Watchlist, error)
}

type MovieRepo struct {
	db      *mongo.Database
	service entity.Services
	resty   *resty.Client
}

func (r *MovieRepo) SearchMovie(ctx context.Context, search entity.SearchRequest) (entity.SearchResponse, error) {
	var res entity.SearchResponse
	resp, err := r.resty.R().
		SetQueryParams(map[string]string{
			"apikey": r.service.OmdbApiKey,
			"s":      url.QueryEscape(search.Query),
			"page":   strconv.Itoa(int(search.Page)),
		}).
		Get(r.service.OmdbHost)

	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *MovieRepo) DetailMovie(ctx context.Context, imdbID string) (entity.DetailResponse, error) {
	var res entity.DetailResponse

	resp, err := r.resty.R().
		SetQueryParams(map[string]string{
			"apikey": r.service.OmdbApiKey,
			"i":      imdbID,
		}).
		Get(r.service.OmdbHost)

	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		return res, err
	}

	if res.Error != "" {
		return res, errors.InvalidArgument(res.Error)
	}

	return res, nil
}

func (r *MovieRepo) AddWatchlistMovie(ctx context.Context, req entity.WatchlistRequest) error {
	movie, err := r.DetailMovie(ctx, req.ImdbID)
	if err != nil {
		return err
	}

	wl := entity.Watchlist{
		UserID:     req.UserID,
		ImdbID:     req.ImdbID,
		MovieTitle: movie.Title,
	}

	opt := &options.UpdateOptions{
		Upsert: aws.Bool(true),
	}

	data := bson.M{"$set": &wl, "$setOnInsert": bson.M{"_id": pkg.NewULID()}}
	filter := bson.M{"user_id": req.UserID, "omdb_id": req.ImdbID}
	_, err = r.db.Collection(m.CollectionWatchlist).UpdateOne(ctx, filter, &data, opt)
	return err
}

func (r *MovieRepo) RemoveWatchlistMovie(ctx context.Context, req entity.WatchlistRequest) error {
	filter := bson.M{"omdb_id": req.ImdbID, "user_id": req.UserID}
	_, err := r.db.Collection(m.CollectionWatchlist).DeleteOne(ctx, filter)
	return err
}

func (r *MovieRepo) GetWatchlist(ctx context.Context, userID string) ([]entity.Watchlist, error) {
	var list []entity.Watchlist

	filter := bson.M{"user_id": userID}
	cursor, err := r.db.Collection(m.CollectionWatchlist).Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	err = cursor.All(ctx, &list)
	if err != nil {
		return list, err
	}

	return list, nil
}

func NewMovieRepo(app *entity.App) IMovieRepo {
	return &MovieRepo{
		db:      app.MongoDbClient,
		service: app.Config.Services,
		resty:   resty.New(),
	}
}
