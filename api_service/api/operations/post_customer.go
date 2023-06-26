// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostCustomerHandlerFunc turns a function with the right signature into a post customer handler
type PostCustomerHandlerFunc func(PostCustomerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostCustomerHandlerFunc) Handle(params PostCustomerParams) middleware.Responder {
	return fn(params)
}

// PostCustomerHandler interface for that can handle valid post customer params
type PostCustomerHandler interface {
	Handle(PostCustomerParams) middleware.Responder
}

// NewPostCustomer creates a new http.Handler for the post customer operation
func NewPostCustomer(ctx *middleware.Context, handler PostCustomerHandler) *PostCustomer {
	return &PostCustomer{Context: ctx, Handler: handler}
}

/* PostCustomer swagger:route POST /customer postCustomer

Create Customer

*/
type PostCustomer struct {
	Context *middleware.Context
	Handler PostCustomerHandler
}

func (o *PostCustomer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostCustomerParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostCustomerBadRequestBody post customer bad request body
//
// swagger:model PostCustomerBadRequestBody
type PostCustomerBadRequestBody struct {

	// error
	Error string `json:"error,omitempty"`
}

// Validate validates this post customer bad request body
func (o *PostCustomerBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post customer bad request body based on context it is used
func (o *PostCustomerBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCustomerBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCustomerBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostCustomerBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostCustomerBody post customer body
//
// swagger:model PostCustomerBody
type PostCustomerBody struct {

	// email
	// Required: true
	Email *string `json:"email"`

	// username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this post customer body
func (o *PostCustomerBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCustomerBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("customer"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	return nil
}

func (o *PostCustomerBody) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("customer"+"."+"username", "body", o.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post customer body based on context it is used
func (o *PostCustomerBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCustomerBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCustomerBody) UnmarshalBinary(b []byte) error {
	var res PostCustomerBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostCustomerOKBody post customer o k body
//
// swagger:model PostCustomerOKBody
type PostCustomerOKBody struct {

	// api key
	APIKey string `json:"api_key,omitempty"`
}

// Validate validates this post customer o k body
func (o *PostCustomerOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post customer o k body based on context it is used
func (o *PostCustomerOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCustomerOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCustomerOKBody) UnmarshalBinary(b []byte) error {
	var res PostCustomerOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostCustomerUnauthorizedBody post customer unauthorized body
//
// swagger:model PostCustomerUnauthorizedBody
type PostCustomerUnauthorizedBody struct {

	// error
	Error string `json:"error,omitempty"`
}

// Validate validates this post customer unauthorized body
func (o *PostCustomerUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post customer unauthorized body based on context it is used
func (o *PostCustomerUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCustomerUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCustomerUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res PostCustomerUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
