package http

import (
	"fmt"
	"log"
	"time"

	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/entity"
	"github.com/labstack/echo/v4"
	"github.com/tylerb/graceful"
)

const defaultPort = "1313"

type Server struct {
	app *entity.App
	e   *echo.Echo
}

func (s *Server) Start() {
	setupDefaultMiddleware(s.e)
	setupRoutes(s.e, s.app)
	ErrorHandler(s.e)

	port := s.app.Config.HttpServerPort
	if port == "" {
		port = defaultPort
	}

	s.e.Server.Addr = fmt.Sprintf(":%v", port)
	s.e.HideBanner = false

	log.Printf("Http server start at port %s", port)
	err := graceful.ListenAndServe(s.e.Server, 5*time.Second)
	if err != nil {
		log.Panic(err)
	}
}

func NewHttpServer(app *entity.App) *Server {
	return &Server{
		app: app,
		e:   echo.New(),
	}
}
