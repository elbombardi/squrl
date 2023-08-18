// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateAccountHandlerFunc turns a function with the right signature into a create account handler
type CreateAccountHandlerFunc func(CreateAccountParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAccountHandlerFunc) Handle(params CreateAccountParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateAccountHandler interface for that can handle valid create account params
type CreateAccountHandler interface {
	Handle(CreateAccountParams, interface{}) middleware.Responder
}

// NewCreateAccount creates a new http.Handler for the create account operation
func NewCreateAccount(ctx *middleware.Context, handler CreateAccountHandler) *CreateAccount {
	return &CreateAccount{Context: ctx, Handler: handler}
}

/*
	CreateAccount swagger:route POST /account accounts createAccount

# Create an account

Create a new account
*/
type CreateAccount struct {
	Context *middleware.Context
	Handler CreateAccountHandler
}

func (o *CreateAccount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewCreateAccountParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
