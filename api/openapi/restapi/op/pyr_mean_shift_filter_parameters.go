// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewPyrMeanShiftFilterParams creates a new PyrMeanShiftFilterParams object
// no default values defined in spec.
func NewPyrMeanShiftFilterParams() PyrMeanShiftFilterParams {

	return PyrMeanShiftFilterParams{}
}

// PyrMeanShiftFilterParams contains all the bound params for the pyr mean shift filter operation
// typically these are obtained from a http.Request
//
// swagger:parameters pyrMeanShiftFilter
type PyrMeanShiftFilterParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	MaxLevel *int64
	/*
	  In: query
	*/
	Picture *string
	/*
	  In: query
	*/
	Sp *float64
	/*
	  In: query
	*/
	Sr *float64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPyrMeanShiftFilterParams() beforehand.
func (o *PyrMeanShiftFilterParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qMaxLevel, qhkMaxLevel, _ := qs.GetOK("maxLevel")
	if err := o.bindMaxLevel(qMaxLevel, qhkMaxLevel, route.Formats); err != nil {
		res = append(res, err)
	}

	qPicture, qhkPicture, _ := qs.GetOK("picture")
	if err := o.bindPicture(qPicture, qhkPicture, route.Formats); err != nil {
		res = append(res, err)
	}

	qSp, qhkSp, _ := qs.GetOK("sp")
	if err := o.bindSp(qSp, qhkSp, route.Formats); err != nil {
		res = append(res, err)
	}

	qSr, qhkSr, _ := qs.GetOK("sr")
	if err := o.bindSr(qSr, qhkSr, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindMaxLevel binds and validates parameter MaxLevel from query.
func (o *PyrMeanShiftFilterParams) bindMaxLevel(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("maxLevel", "query", "int64", raw)
	}
	o.MaxLevel = &value

	return nil
}

// bindPicture binds and validates parameter Picture from query.
func (o *PyrMeanShiftFilterParams) bindPicture(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Picture = &raw

	return nil
}

// bindSp binds and validates parameter Sp from query.
func (o *PyrMeanShiftFilterParams) bindSp(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("sp", "query", "float64", raw)
	}
	o.Sp = &value

	return nil
}

// bindSr binds and validates parameter Sr from query.
func (o *PyrMeanShiftFilterParams) bindSr(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("sr", "query", "float64", raw)
	}
	o.Sr = &value

	return nil
}
