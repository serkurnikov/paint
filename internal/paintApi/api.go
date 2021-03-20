package paintApi

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"paint/internal/app"
	api "paint/pkg/api/proto_files"
	"time"
)

type service struct {
	appl app.Appl
	api.UnimplementedImageProcessingServiceServer
}

func NewServer(appl app.Appl) *grpc.Server {
	srv := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    50 * time.Second,
			Timeout: 10 * time.Second,
		}),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             30 * time.Second,
			PermitWithoutStream: true,
		}),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			unaryServerLogger,
			unaryServerRecover,
			unaryServerAccessLog,
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			streamServerLogger,
			streamServerRecover,
			streamServerAccessLog,
		)),
	)

	api.RegisterImageProcessingServiceServer(srv, &service{
		appl: appl,
	})

	return srv
}
