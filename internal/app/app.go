//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE Appl,Repo

// Package app provides business logic.
package app

import (
	"context"
	"github.com/Jeffail/gabs/v2"
	"paint/internal/apiexternal"
	"paint/pkg/imageProcessingService/colorProcessing/mixcolors"
)

type (
	Ctx = context.Context

	Appl interface {
		ExternalApiTest()
		Render(ctx Ctx)
		BlendColors(colorS1, colorS2 string, numberOfShades int) *gabs.Container
	}

	Repo interface{}

	app struct {
		repo     Repo
		alphaApi apiexternal.Api
	}
)

func (a app) BlendColors(colorS1, colorS2 string, numberOfShades int) *gabs.Container {
	return mixcolors.BlendColors(colorS1, colorS2, numberOfShades)
}

func NewAppl(repo Repo, api apiexternal.Api) Appl {
	return &app{
		repo:     repo,
		alphaApi: api,
	}
}
