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

// PostAccountHandlerFunc turns a function with the right signature into a post account handler
type PostAccountHandlerFunc func(PostAccountParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAccountHandlerFunc) Handle(params PostAccountParams) middleware.Responder {
	return fn(params)
}

// PostAccountHandler interface for that can handle valid post account params
type PostAccountHandler interface {
	Handle(PostAccountParams) middleware.Responder
}

// NewPostAccount creates a new http.Handler for the post account operation
func NewPostAccount(ctx *middleware.Context, handler PostAccountHandler) *PostAccount {
	return &PostAccount{Context: ctx, Handler: handler}
}

/*
	PostAccount swagger:route POST /account postAccount

Create account
*/
type PostAccount struct {
	Context *middleware.Context
	Handler PostAccountHandler
}

func (o *PostAccount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostAccountParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostAccountBadRequestBody post account bad request body
//
// swagger:model PostAccountBadRequestBody
type PostAccountBadRequestBody struct {

	// error
	Error string `json:"error,omitempty"`
}

// Validate validates this post account bad request body
func (o *PostAccountBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post account bad request body based on context it is used
func (o *PostAccountBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAccountBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAccountBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostAccountBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostAccountBody post account body
//
// swagger:model PostAccountBody
type PostAccountBody struct {

	// email
	// Required: true
	Email *string `json:"email"`

	// username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this post account body
func (o *PostAccountBody) Validate(formats strfmt.Registry) error {
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

func (o *PostAccountBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("account"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	return nil
}

func (o *PostAccountBody) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("account"+"."+"username", "body", o.Username); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post account body based on context it is used
func (o *PostAccountBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAccountBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAccountBody) UnmarshalBinary(b []byte) error {
	var res PostAccountBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostAccountInternalServerErrorBody post account internal server error body
//
// swagger:model PostAccountInternalServerErrorBody
type PostAccountInternalServerErrorBody struct {

	// error
	Error string `json:"error,omitempty"`
}

// Validate validates this post account internal server error body
func (o *PostAccountInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post account internal server error body based on context it is used
func (o *PostAccountInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAccountInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAccountInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PostAccountInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostAccountOKBody post account o k body
//
// swagger:model PostAccountOKBody
type PostAccountOKBody struct {

	// api key
	APIKey string `json:"api_key,omitempty"`

	// prefix
	Prefix string `json:"prefix,omitempty"`
}

// Validate validates this post account o k body
func (o *PostAccountOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post account o k body based on context it is used
func (o *PostAccountOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAccountOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAccountOKBody) UnmarshalBinary(b []byte) error {
	var res PostAccountOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostAccountUnauthorizedBody post account unauthorized body
//
// swagger:model PostAccountUnauthorizedBody
type PostAccountUnauthorizedBody struct {

	// error
	Error string `json:"error,omitempty"`
}

// Validate validates this post account unauthorized body
func (o *PostAccountUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post account unauthorized body based on context it is used
func (o *PostAccountUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostAccountUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostAccountUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res PostAccountUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
