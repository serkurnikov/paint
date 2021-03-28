package openapi

import (
	"paint/api/openapi/restapi/op"
	api "paint/api/proto/pb"
)

func (srv *server) HealthCheck(params op.HealthCheckParams) op.HealthCheckResponder {
	ctx, log := fromRequest(params.HTTPRequest)
	status, err := srv.app.HealthCheck(ctx)
	switch {
	default:
		return errHealthCheck(log, err, codeInternal)
	case err == nil:
		return op.NewHealthCheckOK().WithPayload(status)
	}
}

func (srv *server) PyrMeanShiftFilter(params op.PyrMeanShiftFilterParams) op.PyrMeanShiftFilterResponder {
	ctx, log := fromRequest(params.HTTPRequest)
	status, err := srv.app.PyrMeanShiftFilter(ctx, &api.PyrRequest{
		PathPicture: "test",
		Sp:          15,
		Sr:          30,
		MaxLevel:    1,
	})
	switch {
	default:
		return errPyrMeanShiftFilter(log, err, codeInternal)
	case err == nil:
		return op.NewPyrMeanShiftFilterOK().WithPayload(&op.PyrMeanShiftFilterOKBody{Result: status})
	}
}