package grpc

import (
	"log"

	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
)

// Dial grpc server with apm middleware
func Dial(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(apmgrpc.NewUnaryClientInterceptor()))
	if err != nil {
		log.Fatal("could not connect to", addr, err)
	}
	return conn
}
