package openapi

import (
	"paint/api/openapi/restapi/op"
)

func (srv *server) RenderHandlerFunc(params op.RenderParams) op.RenderResponder {
	//ctx, _ := fromRequest(params.HTTPRequest)
	srv.app.UnderPaint(40)

	return op.NewRenderOK().WithPayload(&op.RenderOKBody{Result: "render finished"})
}