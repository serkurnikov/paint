// Example microservice.
package main

import (
	"context"
	"os"
	"os/signal"
	"paint/pkg/def"
	"syscall"
	"time"

	"github.com/powerman/appcfg"
	"github.com/powerman/structlog"
)

var (
	svc = &service{}

	log = structlog.New(structlog.KeyUnit, "main")

	serveStartupTimeout  = appcfg.MustDuration("3s") // must be less than swarm's deploy.update_config.monitor
	serveShutdownTimeout = appcfg.MustDuration("9s") // `docker stop` use 10s between SIGTERM and SIGKILL
)

func main() {
	_ = runServeWithGracefulShutdown()
	select{}
}

func runServeWithGracefulShutdown() error {
	log.Info("started", "version", def.Version())
	defer log.Info("finished", "version", def.Version())

	ctxStartup, cancel := context.WithTimeout(context.Background(), serveStartupTimeout.Value(nil))
	defer cancel()

	ctxShutdown, shutdown := context.WithCancel(context.Background())
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	go func() { <-sigc; shutdown() }()
	go func() {
		<-ctxShutdown.Done()
		time.Sleep(serveShutdownTimeout.Value(nil))
		log.PrintErr("failed to graceful shutdown", "version", def.Version())
		os.Exit(1)
	}()

	return svc.runServe(ctxStartup, ctxShutdown, shutdown)
}
