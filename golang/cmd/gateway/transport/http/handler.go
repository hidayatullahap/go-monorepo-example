package http

import (
	netHttp "net/http"
	"strconv"

	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/action"
	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/entity"
	movieEntity "github.com/hidayatullahap/go-monorepo-example/cmd/movie_service/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg/http"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type Handler struct {
	app    *entity.App
	action action.IGatewayAction
}

func (h *Handler) Register(c echo.Context) error {
	var req entity.RegisterRequest
	err := http.BindAndValidate(c, &req)
	if err != nil {
		return err
	}

	err = h.action.Register(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.NoContent(netHttp.StatusCreated)
}

func (h *Handler) Login(c echo.Context) error {
	var req entity.LoginRequest
	err := http.BindAndValidate(c, &req)
	if err != nil {
		return err
	}

	token, err := h.action.Login(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return c.JSON(netHttp.StatusOK, bson.M{"token": token})
}

func (h *Handler) MovieSearch(c echo.Context) error {
	var result interface{}
	page := c.QueryParam("page")
	pageI, _ := strconv.Atoi(page)

	req := entity.MovieSearchRequest{
		Search: c.QueryParam("search"),
		ImdbID: c.QueryParam("imdb_id"),
		Page:   int64(pageI),
	}

	if req.ImdbID != "" {
		movieDetail, err := h.action.MovieDetail(c.Request().Context(), req)
		if err != nil {
			return err
		}

		result = movieDetail
		return c.JSON(netHttp.StatusOK, result)
	}

	if req.Search != "" {
		movies, err := h.action.MovieSearch(c.Request().Context(), req)
		if err != nil {
			return err
		}

		result = movies
		return c.JSON(netHttp.StatusOK, result)
	}

	return c.NoContent(netHttp.StatusNoContent)
}

func (h *Handler) watchlist(c echo.Context, add bool) error {
	// TODO AUTH for user ID
	var req movieEntity.WatchlistRequest
	req.Fav = add
	req.UserID = "TODO"
	req.ImdbID = c.Param("imdb_id")

	err := h.action.Watchlist(c.Request().Context(), req)

	return err
}

func (h *Handler) GetWatchlist(c echo.Context) error {
	list, err := h.action.UserWatchlist(c.Request().Context(), "TODO")
	if err != nil {
		return err
	}

	data := bson.M{"movies": list}
	return c.JSON(netHttp.StatusOK, data)
}

func (h *Handler) AddToWatchlist(c echo.Context) error {
	return h.watchlist(c, true)
}

func (h *Handler) RemoveFromWatchlist(c echo.Context) error {
	return h.watchlist(c, false)
}

func NewHandler(app *entity.App) *Handler {
	return &Handler{
		app:    app,
		action: action.NewGatewayAction(app),
	}
}
