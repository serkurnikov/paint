package openapi

import (
	"errors"
	oApiErrors "github.com/go-openapi/errors"
	"net/http"
	"paint/internal/app"
)

var errRequireAdmin = errors.New("only admin can make changes")

func (srv *server) authenticate(apiKey string) (*app.Auth, error) {
	switch apiKey {
	case "anonymous":
		return nil, oApiErrors.Unauthenticated("invalid credentials")
	case srv.cfg.APIKeyAdmin:
		return &app.Auth{UserID: "admin"}, nil
	default:
		return &app.Auth{UserID: "user:" + apiKey}, nil
	}
}

func (srv *server) authorize(r *http.Request, principal interface{}) error {
	auth := principal.(*app.Auth)
	if r.Method != "GET" && auth.UserID != "admin" {
		return errRequireAdmin
	}
	return nil
}
