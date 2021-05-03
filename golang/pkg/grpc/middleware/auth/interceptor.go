package auth

import (
	"context"

	"gitlab.com/qasir/web/project/qasircore.git"

	"gitlab.com/qasir/web/project/qasircore.git/grpc/client"
	"gitlab.com/qasir/web/project/qasircore.git/grpc/codes"
	pb "gitlab.com/qasir/web/project/qasircore.git/proto/account/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Info store auth information
type Info struct {
	UserID     uint
	MerchantID uint
}

// UnaryServerInterceptor returns new unary server interceptor that perform per-request auth
func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx, err := authorize(ctx)

	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

func authorize(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Forbidden, "Retrieving metadata is failed")
	}

	authHeader, ok := md["authorization"]
	if !ok {
		return nil, status.Errorf(codes.Unauthorized, "Authorization token is not supplied")
	}

	req := &pb.AuthClientTokenRequest{ClientKey: authHeader[0]}
	res, err := client.AuthClientToken(ctx, req)
	if err != nil {
		return nil, err
	}

	authInfo := Info{
		UserID:     uint(res.UserId),
		MerchantID: uint(res.MerchantId),
	}

	newCtx := context.WithValue(ctx, qasircore.AuthInfo, authInfo)
	return newCtx, nil
}
