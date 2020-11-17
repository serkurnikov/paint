package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/pressly/goose"
)

func connectDB() (*sqlx.DB, error) {
	dbDsnString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "root", "password", "paint",
	)
	dbConn, err := sqlx.Connect("postgres", dbDsnString)
	if err != nil {
		return nil, err
	}

	dbConn.SetMaxIdleConns(50)
	dbConn.SetMaxOpenConns(50)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err = dbConn.PingContext(ctx)
	for err != nil {
		nextErr := dbConn.PingContext(ctx)
		if nextErr == context.DeadlineExceeded {
			return nil, errors.Wrap(err, "connect to postgres")
		}
		err = nextErr
	}

	log.Println("DB connected successful!")
	return dbConn, nil
}

func migrationDB(db *sqlx.DB) error {
	_ = goose.SetDialect("postgres")

	current, err := goose.EnsureDBVersion(db.DB)
	if err != nil {
		return fmt.Errorf("failed to EnsureDBVersion: %v", errors.WithStack(err))
	}

	files, err := ioutil.ReadDir("migrations")
	if err != nil {
		return err
	}

	migrations, err := goose.CollectMigrations("migrations", current, int64(len(files)))
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
