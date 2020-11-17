package openapi

import (
	"encoding/json"
	"net"
	"net/http"
	"paint/api/openapi/model"
	"paint/pkg/def"
	"path"
	"strings"

	"github.com/go-openapi/swag"
	"github.com/powerman/structlog"
	corspkg "github.com/rs/cors"
)

type middlewareFunc func(http.Handler) http.Handler

func noCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Expires", "0")
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		next.ServeHTTP(w, r)
	})
}

// Provide a logger configured using request's context.
//
// Usually it should be one of the first (but after xff, if used) middleware.
func makeLogger(basePath string) middlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := structlog.FromContext(r.Context(), nil)
			log.SetDefaultKeyvals(
				def.LogRemote, r.RemoteAddr,
				def.LogHTTPStatus, "",
				def.LogHTTPMethod, r.Method,
				def.LogFunc, path.Join("/", strings.TrimPrefix(r.URL.Path, basePath)),
			)
			r = r.WithContext(structlog.NewContext(r.Context(), log))

			next.ServeHTTP(w, r)
		})
	}
}

// go-swagger responders panic on error while writing response to client,
// this shouldn't result in crash - unlike a real, reasonable panic.
//
// Usually it should be second middleware (after logger).
func recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panicked := true
		defer func() {
			if panicked {
				const code = http.StatusInternalServerError
				switch err := recover(); err := err.(type) {
				default:
					log := structlog.FromContext(r.Context(), nil)
					log.PrintErr("panic", def.LogHTTPStatus, code, "err", err, structlog.KeyStack, structlog.Auto)
					middlewareError(w, code, "internal error")
				case net.Error:
					log := structlog.FromContext(r.Context(), nil)
					log.PrintErr("recovered", def.LogHTTPStatus, code, "err", err)
					middlewareError(w, code, "internal error")
				}
			}
		}()
		next.ServeHTTP(w, r)
		panicked = false
	})
}

func cors(next http.Handler) http.Handler {
	return corspkg.AllowAll().Handler(next)
}

// MiddlewareError is not a middleware, it's a helper for returning errors
// from middleware.
func middlewareError(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(&model.Error{
		Code:    swag.Int32(int32(code)),
		Message: swag.String(msg),
	})
}
