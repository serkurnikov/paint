package main

import (
	"flag"
	"os"
	"os/signal"
	"paint/internal/dal"
	"paint/internal/def"
	"paint/pkg/flags"
	"paint/pkg/repo"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/powerman/structlog"
)

type (
	dbConfig struct {
		host string
		port int
		user string
		pass string
		name string
	}
)

//nolint:gochecknoglobals
var (
	log    = structlog.New()
	dbConf dbConfig
	cfg    struct {
		version       bool
		logLevel      string
		grpcPort      int
		migrationPath string
	}
)

func Init() {
	def.Init()

	flag.StringVar(&cfg.logLevel, "log.level", "debug", "log level (debug|info|warn|err)")
	flag.IntVar(&cfg.grpcPort, "grpc_service_port", def.GrpcServicePort, "listen on grpc_port (>0)")
	flag.IntVar(&dbConf.port, "db.port", def.DBPort, "db port")
	flag.StringVar(&dbConf.host, "db.host", def.DBHost, "db host")
	flag.StringVar(&dbConf.user, "db.user", def.DBUser, "db user")
	flag.StringVar(&dbConf.name, "db.name", def.DBName, "db name")
	flag.StringVar(&dbConf.pass, "db.pass", def.DBPass, "db pass")
	flag.StringVar(&cfg.migrationPath, "migration_path", def.MigrationPath, "migration path")
}

func main() {
	Init()
	//TODO
	flag.Parse()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)

	switch {
	case cfg.grpcPort <= 0:
		flags.FatalFlagValue("must be > 0", "grpc_service_port", cfg.grpcPort)
	case dbConf.host == "":
		flags.FatalFlagValue("required", "db.host", dbConf.host)
	case dbConf.port <= 0:
		flags.FatalFlagValue("must be > 0", "db.port", dbConf.port)
	case dbConf.user == "":
		flags.FatalFlagValue("required", "db.user", dbConf.user)
	case dbConf.pass == "":
		flags.FatalFlagValue("required", "db.pass", dbConf.pass)
	case dbConf.name == "":
		flags.FatalFlagValue("required", "db.name", dbConf.name)
	case cfg.migrationPath == "":
		flags.FatalFlagValue("required", "migration_path", cfg.migrationPath)
	}

	structlog.DefaultLogger.SetLogLevel(structlog.ParseLevel(cfg.logLevel))

	r, err := dal.New(&repo.Config{
		Host:          dbConf.host,
		Port:          dbConf.port,
		Name:          dbConf.name,
		User:          dbConf.user,
		Pass:          dbConf.pass,
		MaxIdleConns:  100,
		MaxOpenConns:  50,
		MigrationPath: cfg.migrationPath,
	}, log)
	if err != nil {
		log.Fatal(err)
	}

	shutdown, err := RunService(r, log)
	if err != nil {
		log.Fatal(err)
	}

	<-done
	shutdown()
	os.Exit(0)
}
