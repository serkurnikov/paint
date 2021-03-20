package openapi

import (
	"paint/api/openapi/restapi/op"
)

func (srv *server) RenderHandlerFunc(params op.RenderParams) op.RenderResponder {
	ctx, _ := fromRequest(params.HTTPRequest)
	srv.app.Render(ctx)

	//work := highload.Job{Payload: struct{}{}}
	//highload.JobQueue <- work

	return op.NewRenderOK().WithPayload(&op.RenderOKBody{Result: "render finished"})
}