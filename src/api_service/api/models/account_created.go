// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AccountCreated account created
//
// swagger:model AccountCreated
type AccountCreated struct {

	// api key
	APIKey string `json:"api_key,omitempty"`

	// prefix
	Prefix string `json:"prefix,omitempty"`
}

// Validate validates this account created
func (m *AccountCreated) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this account created based on context it is used
func (m *AccountCreated) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AccountCreated) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountCreated) UnmarshalBinary(b []byte) error {
	var res AccountCreated
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
