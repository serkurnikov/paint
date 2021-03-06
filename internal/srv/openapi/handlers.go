package openapi

import (
	"paint/api/openapi/restapi/op"
)

func (srv *server) RenderHandlerFunc(params op.RenderParams) op.RenderResponder {
	ctx, _ := fromRequest(params.HTTPRequest)
	srv.app.Render(ctx)

	return op.NewRenderOK().WithPayload(&op.RenderOKBody{Result: "render finished"})
}

func (srv *server) ScobelHandlerFunc(params op.ScobelParams) op.ScobelResponder {
	srv.app.Scobel()
	return op.NewScobelOK().WithPayload(&op.ScobelOKBody{Result: "scobel finished"})
}