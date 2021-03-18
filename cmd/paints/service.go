package main

import (
	"context"
	"github.com/powerman/structlog"
	"paint/api/openapi/restapi"
	"paint/internal/apiexternal"
	"paint/internal/app"
	"paint/internal/config"
	"paint/internal/dal"
	"paint/internal/gRPC/imageProcessingService"
	"paint/internal/srv/openapi"
	"paint/pkg/concurrent"
	"paint/pkg/serve"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

type service struct {
	cfg *config.ServeConfig
	srv *restapi.Server
}

func (s *service) runServe(ctxStartup, ctxShutdown Ctx, shutdown func()) (err error) {
	log := structlog.FromContext(ctxShutdown, nil)
	if s.cfg == nil {
		s.cfg, err = config.GetServe()
	}
	if err != nil {
		return log.Err("failed to get config", "err", err)
	}

	db, err := connectDB()
	if err != nil {
		return log.Err("err", err)
	}

	if err = migrationDB(db); err != nil {
		return log.Err("err", err)
	}
	
	//rbt := rabbitmq.New(s.cfg.RabbitMQ)
	//if err := rbt.Connect(); err != nil {log.Fatalln(err)}

	//processingAMQP := imageProcessing.NewAMQP(s.cfg.ImageProcessingAMQP, rbt)
	//if err := processingAMQP.Setup(); err != nil {log.Fatalln(err)}

	//rabbitImageProcessing := imageProcessing.NewCreate(rbt)


	alphaApi := apiexternal.NewAlphaVantage()
	repo := dal.New(db)
	processingClient := imageProcessingService.NewImageProcessingClient()




	appl := app.NewAppl(repo, alphaApi, processingClient)
	s.srv, err = openapi.NewServer(appl)
	if err != nil {
		return log.Err("failed to openapi.NewServer", "err", err)
	}

	err = concurrent.Serve(ctxShutdown, shutdown,
		s.serveOpenAPI,
	)
	if err != nil {
		return log.Err("failed to serve", "err", err)
	}
	return nil
}

func (s *service) serveOpenAPI(ctx Ctx) error {
	return serve.OpenAPI(ctx, s.srv, "OpenAPI")
}
