// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"paint/api/openapi/model"
)

// PyrMeanShiftFilterOKCode is the HTTP code returned for type PyrMeanShiftFilterOK
const PyrMeanShiftFilterOKCode int = 200

/*PyrMeanShiftFilterOK response

swagger:response pyrMeanShiftFilterOK
*/
type PyrMeanShiftFilterOK struct {

	/*
	  In: Body
	*/
	Payload *PyrMeanShiftFilterOKBody `json:"body,omitempty"`
}

// NewPyrMeanShiftFilterOK creates PyrMeanShiftFilterOK with default headers values
func NewPyrMeanShiftFilterOK() *PyrMeanShiftFilterOK {

	return &PyrMeanShiftFilterOK{}
}

// WithPayload adds the payload to the pyr mean shift filter o k response
func (o *PyrMeanShiftFilterOK) WithPayload(payload *PyrMeanShiftFilterOKBody) *PyrMeanShiftFilterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pyr mean shift filter o k response
func (o *PyrMeanShiftFilterOK) SetPayload(payload *PyrMeanShiftFilterOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PyrMeanShiftFilterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *PyrMeanShiftFilterOK) PyrMeanShiftFilterResponder() {}

/*PyrMeanShiftFilterDefault General errors using same model as used by go-swagger for validation errors.

swagger:response pyrMeanShiftFilterDefault
*/
type PyrMeanShiftFilterDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *model.Error `json:"body,omitempty"`
}

// NewPyrMeanShiftFilterDefault creates PyrMeanShiftFilterDefault with default headers values
func NewPyrMeanShiftFilterDefault(code int) *PyrMeanShiftFilterDefault {
	if code <= 0 {
		code = 500
	}

	return &PyrMeanShiftFilterDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the pyr mean shift filter default response
func (o *PyrMeanShiftFilterDefault) WithStatusCode(code int) *PyrMeanShiftFilterDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the pyr mean shift filter default response
func (o *PyrMeanShiftFilterDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the pyr mean shift filter default response
func (o *PyrMeanShiftFilterDefault) WithPayload(payload *model.Error) *PyrMeanShiftFilterDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pyr mean shift filter default response
func (o *PyrMeanShiftFilterDefault) SetPayload(payload *model.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PyrMeanShiftFilterDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

func (o *PyrMeanShiftFilterDefault) PyrMeanShiftFilterResponder() {}

type PyrMeanShiftFilterNotImplementedResponder struct {
	middleware.Responder
}

func (*PyrMeanShiftFilterNotImplementedResponder) PyrMeanShiftFilterResponder() {}

func PyrMeanShiftFilterNotImplemented() PyrMeanShiftFilterResponder {
	return &PyrMeanShiftFilterNotImplementedResponder{
		middleware.NotImplemented(
			"operation authentication.PyrMeanShiftFilter has not yet been implemented",
		),
	}
}

type PyrMeanShiftFilterResponder interface {
	middleware.Responder
	PyrMeanShiftFilterResponder()
}
