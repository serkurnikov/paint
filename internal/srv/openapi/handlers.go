package openapi

import (
	"paint/api/openapi/restapi/op"
)

func (srv *server) TestHandlerFunc(params op.TestParams) op.TestResponder {
	//ctx, _ := fromRequest(params.HTTPRequest)
	result := srv.app.UnderPaint(10)

	return op.NewTestOK().WithPayload(&op.TestOKBody{Result: result})
}