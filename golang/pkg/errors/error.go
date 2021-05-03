package errors

import (
	"github.com/hidayatullahap/go-monorepo-example/pkg/grpc/codes"
	"google.golang.org/grpc/status"
)

func InternalError(message string) error {
	return status.Errorf(codes.InternalError, message)
}
