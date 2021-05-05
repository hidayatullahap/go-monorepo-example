package http

import (
	netHttp "net/http"
	"strconv"

	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/action"
	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/entity"
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
		OmdbID: c.QueryParam("omdb_id"),
		Page:   int64(pageI),
	}

	if req.OmdbID != "" {

	} else {
		movies, err := h.action.MovieSearch(c.Request().Context(), req)
		if err != nil {
			return err
		}

		result = movies
	}

	return c.JSON(netHttp.StatusOK, result)
}

func NewHandler(app *entity.App) *Handler {
	return &Handler{
		app:    app,
		action: action.NewGatewayAction(app),
	}
}
