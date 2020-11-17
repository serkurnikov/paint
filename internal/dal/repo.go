package dal

import (
	"context"
	"paint/internal/app"

	"github.com/jmoiron/sqlx"
)

type Ctx = context.Context

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) app.Repo {
	return &repo{db: db}
}
