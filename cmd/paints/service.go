package main

import (
	"context"
	"github.com/powerman/structlog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
	"paint/api/openapi/restapi"
	"paint/internal/apiexternal"
	"paint/internal/app"
	"paint/internal/config"
	"paint/internal/dal"
	"paint/internal/gRPC/imageProcessingService"
	"paint/internal/srv/openapi"
	"paint/migrations"
	"paint/pkg/cobrax"
	"paint/pkg/concurrent"
	"paint/pkg/def"
	"paint/pkg/rabbitmq"
	"paint/pkg/serve"
	"regexp"
)

type Ctx = context.Context

var reg = prometheus.NewPedanticRegistry()

type service struct {
	cfg  *config.ServeConfig
	repo *dal.Repo
	appl *app.App
	srv  *restapi.Server
	rabbitMQ *rabbitmq.RabbitMQ
}

func initService(cmd, serveCmd *cobra.Command) error {
	namespace := regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(def.ProgName, "_")
	initMetrics(reg, namespace)
	dal.InitMetrics(reg, namespace)
	app.InitMetrics(reg)
	openapi.InitMetrics(reg, namespace)

	gooseSQLCmd := cobrax.NewGooseSQLCmd(migrations.Goose(), config.GetGooseSQL)
	cmd.AddCommand(gooseSQLCmd)

	return config.Init(config.FlagSets{
		Serve:    serveCmd.Flags(),
		GooseSQL: gooseSQLCmd.Flags(),
	})
}

func (s *service) runServe(ctxStartup, ctxShutdown Ctx, shutdown func()) (err error) {
	log := structlog.FromContext(ctxShutdown, nil)
	if s.cfg == nil {
		s.cfg, err = config.GetServe()
	}
	if err != nil {
		return log.Err("failed to get config", "err", err)
	}

	err = concurrent.Setup(ctxStartup, map[interface{}]concurrent.SetupFunc{
		&s.repo: s.connectRepo, &s.rabbitMQ: s.connectRabbit,
	})
	if err != nil {
		return log.Err("failed to connect", "err", err)
	}

	alphaApi := apiexternal.NewAlphaVantage()
	gRPCImageProcessingClient := imageProcessingService.NewImageProcessingClient()

	if s.appl == nil {
		s.appl = app.New(s.repo, alphaApi, gRPCImageProcessingClient)
	}

	s.srv, err = openapi.NewServer(s.appl, openapi.Config{
		APIKeyAdmin: s.cfg.APIKeyAdmin,
		Addr:        s.cfg.Addr,
	})
	if err != nil {
		return log.Err("failed to openapi.NewServer", "err", err)
	}

	if err != nil {
		return log.Err("failed to openapi.NewServer", "err", err)
	}

	err = concurrent.Serve(ctxShutdown, shutdown,
		s.serveMetrics,
		s.serveOpenAPI,
	)
	if err != nil {
		return log.Err("failed to serve", "err", err)
	}
	return nil
}

func (s *service) connectRepo(ctx Ctx) (interface{}, error) {
	return dal.New(ctx, s.cfg.SQLGooseDir)
}

func (s *service) connectRabbit(ctx Ctx) (interface{}, error) {
	return rabbitmq.New(s.cfg.RabbitMQ), nil
}

func (s *service) serveMetrics(ctx Ctx) error {
	return serve.Metrics(ctx, s.cfg.MetricsAddr, reg)
}

func (s *service) serveOpenAPI(ctx Ctx) error {
	return serve.OpenAPI(ctx, s.srv, "OpenAPI")
}
