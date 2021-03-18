//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE Appl,Repo

// Package app provides business logic.
package app

import (
	"context"
	"paint/internal/apiexternal"
	pb "paint/internal/gRPC/imageProcessingService/service"
)

type (
	Ctx = context.Context

	Appl interface {
		ExternalApiTest()
		Render(ctx Ctx)
		Scobel()
	}

	Repo interface{}

	App struct {
		repo                  Repo
		alphaApi              apiexternal.Api
		imageProcessingClient pb.ImageProcessingServiceClient
	}
)

func NewAppl(repo Repo, api apiexternal.Api, client pb.ImageProcessingServiceClient) Appl {
	return &App{
		repo:                  repo,
		alphaApi:              api,
		imageProcessingClient: client,
	}
}
