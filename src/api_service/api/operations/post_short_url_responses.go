// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostShortURLOKCode is the HTTP code returned for type PostShortURLOK
const PostShortURLOKCode int = 200

/*
PostShortURLOK Success

swagger:response postShortUrlOK
*/
type PostShortURLOK struct {

	/*
	  In: Body
	*/
	Payload *PostShortURLOKBody `json:"body,omitempty"`
}

// NewPostShortURLOK creates PostShortURLOK with default headers values
func NewPostShortURLOK() *PostShortURLOK {

	return &PostShortURLOK{}
}

// WithPayload adds the payload to the post short Url o k response
func (o *PostShortURLOK) WithPayload(payload *PostShortURLOKBody) *PostShortURLOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post short Url o k response
func (o *PostShortURLOK) SetPayload(payload *PostShortURLOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostShortURLOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostShortURLBadRequestCode is the HTTP code returned for type PostShortURLBadRequest
const PostShortURLBadRequestCode int = 400

/*
PostShortURLBadRequest Bad Request

swagger:response postShortUrlBadRequest
*/
type PostShortURLBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PostShortURLBadRequestBody `json:"body,omitempty"`
}

// NewPostShortURLBadRequest creates PostShortURLBadRequest with default headers values
func NewPostShortURLBadRequest() *PostShortURLBadRequest {

	return &PostShortURLBadRequest{}
}

// WithPayload adds the payload to the post short Url bad request response
func (o *PostShortURLBadRequest) WithPayload(payload *PostShortURLBadRequestBody) *PostShortURLBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post short Url bad request response
func (o *PostShortURLBadRequest) SetPayload(payload *PostShortURLBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostShortURLBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostShortURLUnauthorizedCode is the HTTP code returned for type PostShortURLUnauthorized
const PostShortURLUnauthorizedCode int = 401

/*
PostShortURLUnauthorized Unauthorized

swagger:response postShortUrlUnauthorized
*/
type PostShortURLUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *PostShortURLUnauthorizedBody `json:"body,omitempty"`
}

// NewPostShortURLUnauthorized creates PostShortURLUnauthorized with default headers values
func NewPostShortURLUnauthorized() *PostShortURLUnauthorized {

	return &PostShortURLUnauthorized{}
}

// WithPayload adds the payload to the post short Url unauthorized response
func (o *PostShortURLUnauthorized) WithPayload(payload *PostShortURLUnauthorizedBody) *PostShortURLUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post short Url unauthorized response
func (o *PostShortURLUnauthorized) SetPayload(payload *PostShortURLUnauthorizedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostShortURLUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostShortURLInternalServerErrorCode is the HTTP code returned for type PostShortURLInternalServerError
const PostShortURLInternalServerErrorCode int = 500

/*
PostShortURLInternalServerError Internal Server Error

swagger:response postShortUrlInternalServerError
*/
type PostShortURLInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *PostShortURLInternalServerErrorBody `json:"body,omitempty"`
}

// NewPostShortURLInternalServerError creates PostShortURLInternalServerError with default headers values
func NewPostShortURLInternalServerError() *PostShortURLInternalServerError {

	return &PostShortURLInternalServerError{}
}

// WithPayload adds the payload to the post short Url internal server error response
func (o *PostShortURLInternalServerError) WithPayload(payload *PostShortURLInternalServerErrorBody) *PostShortURLInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post short Url internal server error response
func (o *PostShortURLInternalServerError) SetPayload(payload *PostShortURLInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostShortURLInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}