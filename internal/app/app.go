//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE Appl,Repo

// Package app provides business logic.
package app

import (
	"context"
	"errors"
	"paint/internal/apiexternal"
	pb "paint/internal/gRPC/imageProcessingService/service"
)

var (
	ErrExists = errors.New("value already exists")
	ErrRabbitConnect = errors.New("not rabbit connection")
)

type (
	Ctx = context.Context

	// Appl provides application features (use cases) service.
	Appl interface {
		ExternalApiTest()
		Render(ctx Ctx)
		Scobel()
	}

	// Repo provides data storage.
	Repo interface{}

	Auth struct {
		UserID string
	}

	// App implements interface Appl.
	App struct {
		repo                  Repo
		alphaApi              apiexternal.Api
		imageProcessingClient pb.ImageProcessingServiceClient
	}
)

func New(repo Repo, api apiexternal.Api, client pb.ImageProcessingServiceClient) *App {
	return &App{
		repo:                  repo,
		alphaApi:              api,
		imageProcessingClient: client,
	}
}
