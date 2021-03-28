package image_proc

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/powerman/structlog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	api "paint/api/proto/pb"
	"time"
)

type service struct {
	api.UnimplementedImageProcessingServiceServer
}

type (
	Ctx = context.Context
	Log = *structlog.Logger
)

func NewImageProcServer() *grpc.Server {
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

	api.RegisterImageProcessingServiceServer(srv, &service{})
	return srv
}
