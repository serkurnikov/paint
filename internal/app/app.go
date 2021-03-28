//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE Appl,Repo

package app

import (
	"context"
	"github.com/powerman/structlog"
	"paint/internal/imageProcApi"
)

type (
	Ctx = context.Context

	Appl interface {
		imageProcApi.Api
		HealthCheck(ctx Ctx) (interface{}, error)
	}

	Repo interface{}

	Config struct{}

	App struct {
		cfg          Config
		repo         Repo
		imageProcApi imageProcApi.Api
		logger       *structlog.Logger
	}
)

func New(repo Repo, cfg Config, imageProcApi imageProcApi.Api, logger *structlog.Logger) *App {
	a := &App{
		cfg:          cfg,
		repo:         repo,
		imageProcApi: imageProcApi,
		logger:       logger,
	}
	return a
}

func (a App) HealthCheck(ctx Ctx) (interface{}, error) {
	return "OK", nil
}
