package transport

import (
	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/entity"
	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/transport/http"
)

type Transport struct {
	HttpServer *http.Server
}

func NewTransport(app *entity.App) *Transport {
	return &Transport{
		http.NewHttpServer(app),
	}
}
