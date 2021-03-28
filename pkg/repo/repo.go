package repo

import (
	"io/ioutil"

	"github.com/cenkalti/backoff/v4"
	"github.com/jmoiron/sqlx"
	"github.com/powerman/structlog"
	"github.com/pressly/goose"

	_ "github.com/lib/pq"
)

type Repo struct {
	DB     *sqlx.DB
	Logger *structlog.Logger
}

func New(cfg *Config, log *structlog.Logger) (*Repo, error) {
	var err error

	db, err := connectDB(cfg, log)
	if err != nil {
		return nil, err
	}

	if err := migrationDB(db, cfg.MigrationPath); err != nil {
		return nil, err
	}

	return &Repo{DB: db, Logger: log}, nil
}

func connectDB(cfg *Config, log *structlog.Logger) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)

	var pingDB backoff.Operation = func() error {
		err = db.Ping()
		if err != nil {
			log.Println("DB is not ready...backing off...")
			return err
		}
		log.Println("DB is ready!")
		return nil
	}

	err = backoff.Retry(pingDB, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, err
	}

	log.Println("DB connected successful!")
	return db, nil
}

func migrationDB(db *sqlx.DB, migrationPath string) error {
	var err error

	err = goose.SetDialect("postgres")
	if err != nil {
		return err
	}

	current, err := goose.EnsureDBVersion(db.DB)
	if err != nil {
		return err
	}
	files, err := ioutil.ReadDir(migrationPath)
	if err != nil {
		return err
	}

	migrations, err := goose.CollectMigrations(migrationPath, current, int64(len(files)))
	if err != nil {
		return err
	}

	for _, m := range migrations {
		if err := m.Up(db.DB); err != nil {
			return err
		}
	}

	return nil
}
