package client

import (
	"github.com/pkg/errors"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunGrpcClient(
	serverAddr string,
	registerClient func(conn *grpc.ClientConn) interface{},
) (client interface{}, close func() error, err error) {
	if serverAddr == "" {
		return nil, func() error { return nil }, errors.New("empty serverAddr")
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	}

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return registerClient(conn), conn.Close, nil
}
