// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AssertionOptions1 AssertionOptions1
//
// swagger:model AssertionOptions1
type AssertionOptions1 struct {

	// The base64url encoded clientDataJSON returned by the client
	// Required: true
	ClientDataJSON []string `json:"clientDataJSON"`

	// The base64url encoded id returned by the client
	ID string `json:"id,omitempty"`

	// The base64url encoded rawId returned by the client. If res.rawId is missing, res.id will be used instead. If both are missing an error will be thrown.
	RawID string `json:"rawId,omitempty"`

	// type
	Type []string `json:"type"`

	// The base64url encoded userHandle returned by the client. May be null or an empty string.
	UserHandle string `json:"userHandle,omitempty"`
}

// Validate validates this assertion options1
func (m *AssertionOptions1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientDataJSON(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssertionOptions1) validateClientDataJSON(formats strfmt.Registry) error {

	if err := validate.Required("clientDataJSON", "body", m.ClientDataJSON); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssertionOptions1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssertionOptions1) UnmarshalBinary(b []byte) error {
	var res AssertionOptions1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
