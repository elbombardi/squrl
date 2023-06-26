// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PutCustomerOKCode is the HTTP code returned for type PutCustomerOK
const PutCustomerOKCode int = 200

/*PutCustomerOK Success

swagger:response putCustomerOK
*/
type PutCustomerOK struct {

	/*
	  In: Body
	*/
	Payload *PutCustomerOKBody `json:"body,omitempty"`
}

// NewPutCustomerOK creates PutCustomerOK with default headers values
func NewPutCustomerOK() *PutCustomerOK {

	return &PutCustomerOK{}
}

// WithPayload adds the payload to the put customer o k response
func (o *PutCustomerOK) WithPayload(payload *PutCustomerOKBody) *PutCustomerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put customer o k response
func (o *PutCustomerOK) SetPayload(payload *PutCustomerOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutCustomerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutCustomerBadRequestCode is the HTTP code returned for type PutCustomerBadRequest
const PutCustomerBadRequestCode int = 400

/*PutCustomerBadRequest Bad Request

swagger:response putCustomerBadRequest
*/
type PutCustomerBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PutCustomerBadRequestBody `json:"body,omitempty"`
}

// NewPutCustomerBadRequest creates PutCustomerBadRequest with default headers values
func NewPutCustomerBadRequest() *PutCustomerBadRequest {

	return &PutCustomerBadRequest{}
}

// WithPayload adds the payload to the put customer bad request response
func (o *PutCustomerBadRequest) WithPayload(payload *PutCustomerBadRequestBody) *PutCustomerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put customer bad request response
func (o *PutCustomerBadRequest) SetPayload(payload *PutCustomerBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutCustomerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutCustomerUnauthorizedCode is the HTTP code returned for type PutCustomerUnauthorized
const PutCustomerUnauthorizedCode int = 401

/*PutCustomerUnauthorized Unauthorized

swagger:response putCustomerUnauthorized
*/
type PutCustomerUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *PutCustomerUnauthorizedBody `json:"body,omitempty"`
}

// NewPutCustomerUnauthorized creates PutCustomerUnauthorized with default headers values
func NewPutCustomerUnauthorized() *PutCustomerUnauthorized {

	return &PutCustomerUnauthorized{}
}

// WithPayload adds the payload to the put customer unauthorized response
func (o *PutCustomerUnauthorized) WithPayload(payload *PutCustomerUnauthorizedBody) *PutCustomerUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put customer unauthorized response
func (o *PutCustomerUnauthorized) SetPayload(payload *PutCustomerUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutCustomerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutCustomerNotFoundCode is the HTTP code returned for type PutCustomerNotFound
const PutCustomerNotFoundCode int = 404

/*PutCustomerNotFound Not Found

swagger:response putCustomerNotFound
*/
type PutCustomerNotFound struct {

	/*
	  In: Body
	*/
	Payload *PutCustomerNotFoundBody `json:"body,omitempty"`
}

// NewPutCustomerNotFound creates PutCustomerNotFound with default headers values
func NewPutCustomerNotFound() *PutCustomerNotFound {

	return &PutCustomerNotFound{}
}

// WithPayload adds the payload to the put customer not found response
func (o *PutCustomerNotFound) WithPayload(payload *PutCustomerNotFoundBody) *PutCustomerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put customer not found response
func (o *PutCustomerNotFound) SetPayload(payload *PutCustomerNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutCustomerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}