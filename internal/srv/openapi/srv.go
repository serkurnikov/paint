// Package openapi implements OpenAPI server.
package openapi

import (
	"context"
	"fmt"
	"github.com/go-openapi/runtime"
	"net"
	"net/http"
	"paint/api/openapi/restapi"
	"paint/api/openapi/restapi/op"
	"paint/internal/app"
	"paint/pkg/def"
	"paint/pkg/netx"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/powerman/structlog"
	"github.com/sebest/xff"
)

/*var (
	MaxWorker = os.Getenv("MAX_WORKERS")
	MaxQueue  = os.Getenv("MAX_QUEUE")
)*/

type (
	Ctx = context.Context
	Log = *structlog.Logger

	Config struct {
		APIKeyAdmin string
		Addr        netx.Addr
		BasePath    string
	}

	server struct {
		app app.Appl
		cfg Config
	}
)

// NewServer returns OpenAPI server configured to listen on the TCP network
// address cfg.Host:cfg.Port and handle requests on incoming connections.
func NewServer(appl app.Appl, config Config) (*restapi.Server, error) {
	srv := &server{
		app: appl,
		cfg: config,
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, fmt.Errorf("load embedded swagger spec: %w", err)
	}
	if config.BasePath == "" {
		config.BasePath = swaggerSpec.BasePath()
	}
	swaggerSpec.Spec().BasePath = config.BasePath


	api := op.NewPaintAPI(swaggerSpec)
	api.Logger = structlog.New(structlog.KeyUnit, "swagger").Printf
	api.APIKeyAuth = srv.authenticate
	api.APIAuthorizer = runtime.AuthorizerFunc(srv.authorize)

	api.RenderHandler = op.RenderHandlerFunc(srv.RenderHandlerFunc)

	server := restapi.NewServer(api)
	server.Host = "localhost"
	server.Port = 9000

	//dispatcher := highload.NewDispatcher(100)
	//dispatcher.Run()

	// The middleware executes before anything.
	api.UseSwaggerUI()
	globalMiddlewares := func(handler http.Handler) http.Handler {
		xffmw, _ := xff.Default()
		logger := makeLogger(swaggerSpec.BasePath())
		return noCache(xffmw.Handler(logger(recovery(
			middleware.Spec(swaggerSpec.BasePath(), restapi.FlatSwaggerJSON,
				cors(handler))))))
	}
	
	// The middleware executes after serving /swagger.json and routing,
	// but before authentication, binding and validation.
	middlewares := func(handler http.Handler) http.Handler {
		return handler
	}
	server.SetHandler(globalMiddlewares(api.Serve(middlewares)))

	log := structlog.New()
	log.Info("OpenAPI protocol", "version", swaggerSpec.Spec().Info.Version)
	return server, nil
}

func fromRequest(r *http.Request) (Ctx, Log) {
	ctx := r.Context()
	remoteIP, _, _ := net.SplitHostPort(r.RemoteAddr)
	ctx = def.NewContextWithRemoteIP(ctx, remoteIP)
	log := structlog.FromContext(ctx, nil)
	return ctx, log
}
