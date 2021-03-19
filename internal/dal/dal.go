package dal

import (
	"context"
	"paint/internal/app"
	"paint/pkg/repo"
	"time"
)

type Ctx = context.Context

const (
	schemaVersion = 0
	dbMaxOpenCons = 0
	dbMaxIdleCons = 5
)

type Repo struct {
	*repo.Repo
}

func New(ctx Ctx, dir string) (_ *Repo, err error) {
	returnErrs := []error{
		app.ErrExists,
	}

	r := &Repo{}
	r.Repo, err = repo.New(ctx, repo.Config{
		GooseDir:      dir,
		SchemaVersion: schemaVersion,
		Metric:        repo.Metrics{},
		ReturnErrs:    returnErrs,
	})

	if err != nil {
		return nil, err
	}

	r.DB.SetMaxOpenConns(dbMaxOpenCons)
	r.DB.SetMaxIdleConns(dbMaxIdleCons)
	r.SchemaVer.HoldSharedLock(ctx, time.Second)
	return r, nil
}
