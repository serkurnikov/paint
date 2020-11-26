//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock.$GOFILE Appl,Repo

// Package app provides business logic.
package app

import (
	"context"
	"paint/internal/apiexternal"
)

type (
	Ctx = context.Context

	Appl interface {
		ExternalApiTest()
		UnderPaint(tileSize int)
		Scobel()
	}

	Repo interface{}

	App struct {
		repo     Repo
		alphaApi apiexternal.Api
	}
)

func NewAppl(repo Repo, api apiexternal.Api) Appl {
	return &App{
		repo:     repo,
		alphaApi: api,
	}
}
