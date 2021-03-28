package dal

import (
	"github.com/powerman/structlog"
	"paint/pkg/repo"
)

type Repo struct {
	*repo.Repo
}

func New(cfg *repo.Config, logger *structlog.Logger) (*Repo, error) {
	r := &Repo{}
	var err error
	r.Repo, err = repo.New(cfg, logger)
	if err != nil {
		return nil, err
	}
	return r, nil
}
