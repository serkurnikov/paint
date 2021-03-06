// Package serve provides helpers to start and shutdown network services.
package serve

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"paint/pkg/def"
	"paint/pkg/netx"

	"github.com/powerman/must"
	"github.com/powerman/structlog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Ctx is a synonym for convenience.
type Ctx = context.Context

// OpenAPIServer implemented by *restapi.Server generated by go-swagger.
type OpenAPIServer interface {
	HTTPListener() (net.Listener, error)
	TLSListener() (net.Listener, error)
	Serve() error
	Shutdown() error
}

// OpenAPI starts HTTP/HTTPS server using srv logged as imageProcessingService.
// It runs until failed or ctx.Done.
func OpenAPI(ctx Ctx, srv OpenAPIServer, service string) error {
	log := structlog.FromContext(ctx, nil).New(def.LogServer, service)

	for _, f := range []func() (net.Listener, error){srv.HTTPListener, srv.TLSListener} {
		ln, err := f()
		if err != nil {
			return fmt.Errorf("listen: %w", err)
		}
		if ln != nil {
			host, port, err := net.SplitHostPort(ln.Addr().String())
			must.NoErr(err)
			log.Info("serve", def.LogHost, host, def.LogPort, port)
		}
	}

	go func() { <-ctx.Done(); _ = srv.Shutdown() }()
	err := srv.Serve()
	if err != nil {
		return log.Err("failed to serve", "err", err)
	}
	log.Info("shutdown")
	return nil
}

// HTTP starts HTTP server on addr using handler logged as imageProcessingService.
// It runs until failed or ctx.Done.
func HTTP(ctx Ctx, addr netx.Addr, handler http.Handler, service string) error {
	log := structlog.FromContext(ctx, nil).New(def.LogServer, service)

	srv := &http.Server{
		Addr:    addr.String(),
		Handler: handler,
	}

	log.Info("serve", def.LogHost, addr.Host(), def.LogPort, addr.Port())
	errc := make(chan error, 1)
	go func() { errc <- srv.ListenAndServe() }()

	var err error
	select {
	case err = <-errc:
	case <-ctx.Done():
		err = srv.Shutdown(context.Background())
	}
	if err != nil {
		return log.Err("failed to serve", "err", err)
	}
	log.Info("shutdown")
	return nil
}

// Metrics starts HTTP server on addr path /metrics using reg as
// prometheus handler.
func Metrics(ctx Ctx, addr netx.Addr, reg *prometheus.Registry) error {
	handler := promhttp.InstrumentMetricHandler(reg, promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	mux := http.NewServeMux()
	mux.Handle("/metrics", handler)
	return HTTP(ctx, addr, mux, "Prometheus metrics")
}
