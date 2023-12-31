// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LinkUpdated link updated
//
// swagger:model LinkUpdated
type LinkUpdated struct {

	// long url
	LongURL string `json:"long_url,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// tracking status
	TrackingStatus string `json:"tracking_status,omitempty"`
}

// Validate validates this link updated
func (m *LinkUpdated) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this link updated based on context it is used
func (m *LinkUpdated) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LinkUpdated) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkUpdated) UnmarshalBinary(b []byte) error {
	var res LinkUpdated
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
