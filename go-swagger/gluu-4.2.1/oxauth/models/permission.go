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

// Permission Permission
//
// List of UmaPermission granted to RPT. A permission is (requested or granted) authorized access to a particular resource with some number of scopes  bound to that resource.
//
// swagger:model Permission
type Permission struct {

	// Integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this permission will expire. If the token-level exp value pre-dates a permission-level exp value, the token-level value takes precedence.
	Exp int32 `json:"exp,omitempty"`

	// Integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this permission was originally issued. If the token-level iat value post-dates a permission-level iat value, the token-level value takes precedence.
	Iat int32 `json:"iat,omitempty"`

	// Integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating the time before which this permission is not valid. If the token-level nbf value post-dates a permission-level nbf value, the token-level value takes precedence.
	Nbf int32 `json:"nbf,omitempty"`

	// A string that uniquely identifies the protected resource, access to which has been granted to this client on behalf of this requesting party. The identifier MUST correspond to a resource that was previously registered as protected.
	// Required: true
	ResourceID *string `json:"resource_id"`

	// An array referencing zero or more strings representing scopes to which access was granted for this resource. Each string MUST correspond to a scope that was registered by this resource server for the referenced resource.
	// Required: true
	ResourceScopes []string `json:"resource_scopes"`
}

// Validate validates this permission
func (m *Permission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceScopes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Permission) validateResourceID(formats strfmt.Registry) error {

	if err := validate.Required("resource_id", "body", m.ResourceID); err != nil {
		return err
	}

	return nil
}

func (m *Permission) validateResourceScopes(formats strfmt.Registry) error {

	if err := validate.Required("resource_scopes", "body", m.ResourceScopes); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Permission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Permission) UnmarshalBinary(b []byte) error {
	var res Permission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
