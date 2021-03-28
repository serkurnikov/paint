package main

import (
	"context"
	"github.com/powerman/structlog"
	"github.com/prometheus/client_golang/prometheus"
	"paint/api/openapi/restapi"
	"paint/configs"
	"paint/internal/app"
	"paint/internal/dal"
	"paint/internal/srv/openapi"
	"paint/pkg/concurrent"
	"paint/pkg/netx"
	"paint/pkg/serve"
)

type Ctx = context.Context

var reg = prometheus.NewPedanticRegistry()

type service struct {
	cfg  *configs.Config
	repo *dal.Repo
	appl *app.App
	srv  *restapi.Server
}

func (s *service) runServe(ctxStartup, ctxShutdown Ctx, shutdown func()) (err error) {
	logger := structlog.FromContext(ctxShutdown, nil)

	if s.cfg == nil {
		s.cfg, _ = configs.Init()
	}
	err = concurrent.Setup(ctxStartup, map[interface{}]concurrent.SetupFunc{
		&s.repo: s.connectRepo,
	})
	if s.appl == nil {
		s.appl = app.New(s.repo, app.Config{}, nil, logger)
	}
	s.srv, err = openapi.NewServer(s.appl, openapi.Config{
		Addr: netx.NewAddr(s.cfg.Server.Host, s.cfg.Server.Port),
	})
	if err != nil {
		return logger.Err("failed to openapi.NewServer", "err", err)
	}

	err = concurrent.Serve(ctxShutdown, shutdown,
		s.serveMetrics,
		s.serveOpenAPI,
	)
	if err != nil {
		return logger.Err("failed to serve", "err", err)
	}
	return nil
}

func (s *service) connectRepo(ctx Ctx) (interface{}, error) {
	return dal.New(s.cfg.GetDbConfig(), structlog.FromContext(ctx, nil))
}

func (s *service) serveMetrics(ctx Ctx) error {
	addr := netx.NewAddr(s.cfg.Server.Host, s.cfg.Server.MetricAddrPort)
	return serve.Metrics(ctx, addr, reg)
}

func (s *service) serveOpenAPI(ctx Ctx) error {
	return serve.OpenAPI(ctx, s.srv, "OpenAPI")
}
