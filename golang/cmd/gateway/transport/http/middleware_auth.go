package http

import (
	"context"
	"strings"

	"github.com/hidayatullahap/go-monorepo-example/cmd/gateway/entity"
	"github.com/hidayatullahap/go-monorepo-example/pkg/auth"
	"github.com/hidayatullahap/go-monorepo-example/pkg/grpc"
	"github.com/hidayatullahap/go-monorepo-example/pkg/grpc/codes"
	"github.com/hidayatullahap/go-monorepo-example/pkg/http"
	pb "github.com/hidayatullahap/go-monorepo-example/pkg/proto/auth"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/status"
)

type MiddlewareAuth struct {
	config entity.Config
}

func (h *MiddlewareAuth) AuthCheckToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		if ctx == nil {
			ctx = context.Background()
		}

		conn, err := grpc.Dial(h.config.Services.AuthHost)
		if err != nil {
			return err
		}

		defer conn.Close()

		headerToken := c.Request().Header.Get(echo.HeaderAuthorization)
		token := strings.Replace(headerToken, "Bearer ", "", -1)
		pbReq := &pb.Token{
			Token: token,
		}

		user, err := pb.NewAuthClient(conn).Auth(ctx, pbReq)
		if err != nil {
			st, _ := status.FromError(err)
			resp := http.Response{
				Code:    st.Code(),
				Message: codes.StatusMessage[st.Code()],
				Errors:  []string{st.Message()},
			}
			return resp.JSON(c)
		}

		var payload auth.TokenPayload
		payload.UserID = user.Id
		payload.Username = user.Username

		c.Set(auth.ContextTokenValue, payload)
		return next(c)
	}
}

func NewMiddlewareAuth(config entity.Config) *MiddlewareAuth {
	return &MiddlewareAuth{config}
}
