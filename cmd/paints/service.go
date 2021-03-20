package main

import (
	"context"
	"paint/internal/app"
	"paint/internal/dal"
	"paint/internal/paintApi"
	"paint/pkg/concurrent"
	"paint/pkg/netx"
	"paint/pkg/serve"

	"github.com/pkg/errors"
	"github.com/powerman/structlog"
	"google.golang.org/grpc"
)

type Ctx = context.Context

type service struct {
	srv *grpc.Server
}

func RunService(repo *dal.Repo, log *structlog.Logger) (context.CancelFunc, error) {
	appl := app.NewAppl(repo, nil)
	server := paintApi.NewServer(appl)
	s := service{
		srv: server,
	}

	ctxShutdown, shutdown := context.WithCancel(context.Background())
	err := concurrent.Serve(ctxShutdown, shutdown,
		s.serveGRPC,
	)
	if err != nil {
		return nil, errors.Wrap(err, "starting serve services")
	}

	return shutdown, nil
}

func (s *service) serveGRPC(ctx Ctx) error {
	addr := netx.NewAddr("", cfg.grpcPort)
	return serve.ServerGRPC(ctx, addr, s.srv)
}