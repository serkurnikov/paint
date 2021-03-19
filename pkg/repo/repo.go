package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/powerman/narada4d/schemaver"
	"github.com/powerman/structlog"
	"github.com/pressly/goose"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/powerman/go-service-example/pkg/reflectx"
)

type Ctx = context.Context

var (
	ErrSchemaVer = errors.New("unsupported DB schema version")
)

const (
	DbHost     = "localhost"
	DbPort     = 5432
	DbUser     = "root"
	DbPassword = "password"
	DbName     = "paint"

	DefaultMaxIdleCons = 50
	DefaultMaxOpenCons  = 50

	DefaultDialect = "postgres"
)

type Config struct {
	GooseDir      string
	SchemaVersion int64
	Metric        Metrics
	ReturnErrs    []error
}

type Repo struct {
	DB            *sqlx.DB
	SchemaVer     *schemaver.SchemaVer
	schemaVersion string
	returnErrs    []error
	metric        Metrics
	log           *structlog.Logger
}

func New(ctx Ctx, cfg Config) (*Repo, error) {
	log := structlog.FromContext(ctx, nil)

	dbDsnString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DbHost, DbPort, DbUser, DbPassword, DbName,
	)

	dbConn, err := sqlx.Connect(DefaultDialect, dbDsnString)
	if err != nil {
		return nil, err
	}

	dbConn.SetMaxIdleConns(DefaultMaxIdleCons)
	dbConn.SetMaxOpenConns(DefaultMaxOpenCons)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err = dbConn.PingContext(ctx)
	for err != nil {
		nextErr := dbConn.PingContext(ctx)
		if nextErr == context.DeadlineExceeded {
			log.WarnIfFail(dbConn.Close)
			return nil, errors.Wrap(err, "connect to postgres")
		}
		err = nextErr
	}

	if err = migration(dbConn); err != nil {
		return nil, errors.Wrap(err, "migrate postgres")
	}

	r := &Repo{
		DB:            dbConn,
		schemaVersion: strconv.Itoa(int(cfg.SchemaVersion)),
		returnErrs:    cfg.ReturnErrs,
		metric:        cfg.Metric,
		log:           log,
	}
	return r, nil
}

func migration(db *sqlx.DB) error {
	_ = goose.SetDialect(DefaultDialect)

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

func (r *Repo) Close() {
	r.log.WarnIfFail(r.DB.Close)
	r.log.WarnIfFail(r.SchemaVer.Close)
}

// Turn sqlx errors like `missing destination …` into panics
// https://github.com/jmoiron/sqlx/issues/529. As we can't distinguish
// between sqlx and other errors except driver ones, let's hope filtering
// driver errors is enough and there are no other non-driver regular errors.
func (r *Repo) strict(err error) error {
	switch {
	case err == nil:
	case errors.As(err, new(*mysql.MySQLError)):
	case errors.Is(err, ErrSchemaVer):
	case errors.Is(err, sql.ErrNoRows):
	case errors.Is(err, context.Canceled):
	case errors.Is(err, context.DeadlineExceeded):
	default:
		for i := range r.returnErrs {
			if errors.Is(err, r.returnErrs[i]) {
				return err
			}
		}
		panic(err)
	}
	return err
}

func (r *Repo) schemaLock(f func() error) error {
	ver := r.SchemaVer.SharedLock()
	defer r.SchemaVer.Unlock()
	if ver != r.schemaVersion {
		return fmt.Errorf("schema version %s, need %s: %w", ver, r.schemaVersion, ErrSchemaVer)
	}
	return f()
}

// NoTx provides DAL method wrapper with:
// - converting sqlx errors which are actually bugs into panics,
// - ensure valid schema version while accessing DB,
// - general metrics for DAL methods,
// - wrapping errors with DAL method name.
func (r *Repo) NoTx(f func() error) (err error) {
	methodName := reflectx.CallerMethodName(1)
	return r.strict(r.schemaLock(r.metric.instrument(methodName, func() error {
		err := f()
		if err != nil {
			err = fmt.Errorf("%s: %w", methodName, err)
		}
		return err
	})))
}

// Tx provides DAL method wrapper with:
// - converting sqlx errors which are actually bugs into panics,
// - ensure valid schema version while accessing DB,
// - general metrics for DAL methods,
// - wrapping errors with DAL method name,
// - transaction.
func (r *Repo) Tx(ctx Ctx, opts *sql.TxOptions, f func(*sqlx.Tx) error) (err error) {
	methodName := reflectx.CallerMethodName(1)
	return r.strict(r.schemaLock(r.metric.instrument(methodName, func() error {
		tx, err := r.DB.BeginTxx(ctx, opts)
		if err == nil {
			defer func() {
				if err := recover(); err != nil {
					if err := tx.Rollback(); err != nil {
						log := structlog.FromContext(ctx, nil)
						log.Warn("failed to tx.Rollback", "method", methodName, "err", err)
					}
					panic(err)
				}
			}()
			err = f(tx)
			if err == nil {
				err = tx.Commit()
			} else if err := tx.Rollback(); err != nil {
				log := structlog.FromContext(ctx, nil)
				log.Warn("failed to tx.Rollback", "method", methodName, "err", err)
			}
		}
		if err != nil {
			err = fmt.Errorf("%s: %w", methodName, err)
		}
		return err
	})))
}
