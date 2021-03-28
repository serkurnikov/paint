package main

import (
	"github.com/powerman/structlog"
	"os"
	"os/signal"
	"paint/configs"
	"paint/pkg/def"
	"paint/pkg/flags"
	"syscall"
)

func main() {
	def.Init()
	cfg, _ := configs.Init()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)

	switch {
	case cfg.ImageProcServer.Port <= 0:
		flags.FatalFlagValue("must be > 0", "grpc_service_port", cfg.ImageProcServer.Port)
	case cfg.DataBase.Host == "":
		flags.FatalFlagValue("required", "db.host", cfg.DataBase.Host)
	case cfg.DataBase.Port <= 0:
		flags.FatalFlagValue("must be > 0", "db.port", cfg.DataBase.Port)
	case cfg.DataBase.User == "":
		flags.FatalFlagValue("required", "db.user", cfg.DataBase.User)
	case cfg.DataBase.Pass == "":
		flags.FatalFlagValue("required", "db.pass", cfg.DataBase.Pass)
	case cfg.DataBase.Name == "":
		flags.FatalFlagValue("required", "db.name", cfg.DataBase.Name)
	case cfg.DataBase.MigrationPath == "":
		flags.FatalFlagValue("required", "migration_path", cfg.DataBase.MigrationPath)
	}

	logger := structlog.DefaultLogger.SetLogLevel(structlog.ParseLevel(cfg.ImageProcServer.LogLevel))

	shutdown, err := RunService(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}

	<-done
	shutdown()
	os.Exit(0)
}
