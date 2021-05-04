package errors

import (
	"errors"

	"github.com/hidayatullahap/go-monorepo-example/pkg/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNotFound = errors.New("resource not found")
)

func InternalError(message string) error {
	return status.Errorf(codes.InternalError, message)
}

func InvalidArgument(message string) error {
	return status.Errorf(codes.InvalidArgument, message)
}

func NotFound(message string) error {
	return status.Errorf(codes.NotFound, message)
}
