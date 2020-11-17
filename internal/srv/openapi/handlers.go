package openapi

import (
	"paint/api/openapi/restapi/op"
)

func (srv *server) TestHandlerFunc(params op.TestParams) op.TestResponder {
	///ctx, _ := fromRequest(params.HTTPRequest)

	return op.NewTestOK().WithPayload(&op.TestOKBody{Result: "test finish"})
}