//go:generate gobin -m -run github.com/cheekybits/genny -in=$GOFILE -out=gen.$GOFILE gen "HealthCheck=PyrMeanShiftFilter"
//go:generate sed -i -e "\\,^//go:generate,d" gen.$GOFILE

package openapi

import (
	"github.com/go-openapi/swag"
	"net/http"
	"paint/api/openapi/model"
	"paint/api/openapi/restapi/op"
	image_proc "paint/internal/srv/grpc/image-proc"
	"paint/pkg/def"
)

func errHealthCheck(log image_proc.Log, err error, code errCode) op.HealthCheckResponder {
	return op.NewHealthCheckDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(code.extra),
		Message: swag.String(logger(log, err, code)),
	})
}

func errPyrMeanShiftFilter(log image_proc.Log, err error, code errCode) op.PyrMeanShiftFilterResponder {
	return op.NewPyrMeanShiftFilterDefault(code.status).WithPayload(&model.Error{
		Code:    swag.Int32(code.extra),
		Message: swag.String(logger(log, err, code)),
	})
}

func logger(log image_proc.Log, err error, code errCode) string {
	if code.status < http.StatusInternalServerError {
		log.Info("client error", def.LogHTTPStatus, code.status, "code", code.extra, "err", err)
	} else {
		log.PrintErr("server error", def.LogHTTPStatus, code.status, "code", code.extra, "err", err)
	}

	msg := err.Error()
	if code.status == http.StatusInternalServerError {
		msg = "internal error"
	}
	return msg
}