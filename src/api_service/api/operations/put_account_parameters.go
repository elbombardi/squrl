// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewPutAccountParams creates a new PutAccountParams object
//
// There are no default values defined in the spec.
func NewPutAccountParams() PutAccountParams {

	return PutAccountParams{}
}

// PutAccountParams contains all the bound params for the put account operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutAccount
type PutAccountParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The admin API key.
	  Required: true
	  In: header
	*/
	XAPIKEY string
	/*
	  Required: true
	  In: body
	*/
	Body PutAccountBody
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutAccountParams() beforehand.
func (o *PutAccountParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindXAPIKEY(r.Header[http.CanonicalHeaderKey("X-API-KEY")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body PutAccountBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body", ""))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXAPIKEY binds and validates parameter XAPIKEY from header.
func (o *PutAccountParams) bindXAPIKEY(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("X-API-KEY", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("X-API-KEY", "header", raw); err != nil {
		return err
	}
	o.XAPIKEY = raw

	return nil
}