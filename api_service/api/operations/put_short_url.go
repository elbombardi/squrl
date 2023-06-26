// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PutShortURLHandlerFunc turns a function with the right signature into a put short URL handler
type PutShortURLHandlerFunc func(PutShortURLParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutShortURLHandlerFunc) Handle(params PutShortURLParams) middleware.Responder {
	return fn(params)
}

// PutShortURLHandler interface for that can handle valid put short URL params
type PutShortURLHandler interface {
	Handle(PutShortURLParams) middleware.Responder
}

// NewPutShortURL creates a new http.Handler for the put short URL operation
func NewPutShortURL(ctx *middleware.Context, handler PutShortURLHandler) *PutShortURL {
	return &PutShortURL{Context: ctx, Handler: handler}
}

/* PutShortURL swagger:route PUT /short-url putShortUrl

Update ShortURL

*/
type PutShortURL struct {
	Context *middleware.Context
	Handler PutShortURLHandler
}

func (o *PutShortURL) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutShortURLParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PutShortURLBadRequestBody put short URL bad request body
//
// swagger:model PutShortURLBadRequestBody
type PutShortURLBadRequestBody struct {

	// error
	Error string `json:"error,omitempty"`
}

// Validate validates this put short URL bad request body
func (o *PutShortURLBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put short URL bad request body based on context it is used
func (o *PutShortURLBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutShortURLBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutShortURLBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PutShortURLBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutShortURLBody put short URL body
//
// swagger:model PutShortURLBody
type PutShortURLBody struct {

	// new long url
	NewLongURL string `json:"new_long_url,omitempty"`

	// short url key
	// Required: true
	ShortURLKey *string `json:"short_url_key"`

	// status
	// Enum: [active inactive]
	Status string `json:"status,omitempty"`

	// tracking status
	// Enum: [active inactive]
	TrackingStatus string `json:"tracking_status,omitempty"`
}

// Validate validates this put short URL body
func (o *PutShortURLBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateShortURLKey(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTrackingStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutShortURLBody) validateShortURLKey(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"short_url_key", "body", o.ShortURLKey); err != nil {
		return err
	}

	return nil
}

var putShortUrlBodyTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		putShortUrlBodyTypeStatusPropEnum = append(putShortUrlBodyTypeStatusPropEnum, v)
	}
}

const (

	// PutShortURLBodyStatusActive captures enum value "active"
	PutShortURLBodyStatusActive string = "active"

	// PutShortURLBodyStatusInactive captures enum value "inactive"
	PutShortURLBodyStatusInactive string = "inactive"
)

// prop value enum
func (o *PutShortURLBody) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, putShortUrlBodyTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PutShortURLBody) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("body"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

var putShortUrlBodyTypeTrackingStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		putShortUrlBodyTypeTrackingStatusPropEnum = append(putShortUrlBodyTypeTrackingStatusPropEnum, v)
	}
}

const (

	// PutShortURLBodyTrackingStatusActive captures enum value "active"
	PutShortURLBodyTrackingStatusActive string = "active"

	// PutShortURLBodyTrackingStatusInactive captures enum value "inactive"
	PutShortURLBodyTrackingStatusInactive string = "inactive"
)

// prop value enum
func (o *PutShortURLBody) validateTrackingStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, putShortUrlBodyTypeTrackingStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PutShortURLBody) validateTrackingStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.TrackingStatus) { // not required
		return nil
	}

	// value enum
	if err := o.validateTrackingStatusEnum("body"+"."+"tracking_status", "body", o.TrackingStatus); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this put short URL body based on context it is used
func (o *PutShortURLBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutShortURLBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutShortURLBody) UnmarshalBinary(b []byte) error {
	var res PutShortURLBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutShortURLNotFoundBody put short URL not found body
//
// swagger:model PutShortURLNotFoundBody
type PutShortURLNotFoundBody struct {

	// error
	Error string `json:"error,omitempty"`
}

// Validate validates this put short URL not found body
func (o *PutShortURLNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put short URL not found body based on context it is used
func (o *PutShortURLNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutShortURLNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutShortURLNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PutShortURLNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutShortURLOKBody put short URL o k body
//
// swagger:model PutShortURLOKBody
type PutShortURLOKBody struct {

	// long url
	LongURL string `json:"long_url,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// tracking status
	TrackingStatus string `json:"tracking_status,omitempty"`
}

// Validate validates this put short URL o k body
func (o *PutShortURLOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put short URL o k body based on context it is used
func (o *PutShortURLOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutShortURLOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutShortURLOKBody) UnmarshalBinary(b []byte) error {
	var res PutShortURLOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutShortURLUnauthorizedBody put short URL unauthorized body
//
// swagger:model PutShortURLUnauthorizedBody
type PutShortURLUnauthorizedBody struct {

	// error
	Error string `json:"error,omitempty"`
}

// Validate validates this put short URL unauthorized body
func (o *PutShortURLUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put short URL unauthorized body based on context it is used
func (o *PutShortURLUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutShortURLUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutShortURLUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res PutShortURLUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}