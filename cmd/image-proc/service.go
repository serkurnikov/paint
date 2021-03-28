package main

import (
	"context"
	"paint/configs"
	"paint/internal/dal"
	image_proc "paint/internal/srv/grpc/image-proc"
	"paint/pkg/concurrent"
	"paint/pkg/netx"
	"paint/pkg/serve"

	"github.com/pkg/errors"
	"github.com/powerman/structlog"
	"google.golang.org/grpc"
)

type Ctx = context.Context

type service struct {
	cfg    *configs.Config
	repo   *dal.Repo
	srv    *grpc.Server
	logger *structlog.Logger
}

func RunService(cfg *configs.Config, log *structlog.Logger) (context.CancelFunc, error) {

	server := image_proc.NewImageProcServer()
	s := service{
		cfg:    cfg,
		srv:    server,
		logger: log,
	}

	repoErrors := concurrent.Setup(context.Background(), map[interface{}]concurrent.SetupFunc{
		&s.repo: s.connectRepo,
	})
	if repoErrors != nil {
		return nil, errors.Wrap(repoErrors, "starting database services")
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
	addr := netx.NewAddr(s.cfg.ImageProcServer.Host, s.cfg.ImageProcServer.Port)
	return serve.ServerGRPC(ctx, addr, s.srv)
}

func (s *service) connectRepo(ctx Ctx) (interface{}, error) {
	return dal.New(s.cfg.GetDbConfig(), structlog.FromContext(ctx, nil))
}
