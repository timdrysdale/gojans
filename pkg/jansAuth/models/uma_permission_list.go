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

// UMAPermissionList UmaPermissionList
//
// swagger:model UmaPermissionList
type UMAPermissionList struct {

	// Number of seconds since January 1 1970 UTC, indicating when this token will expire.
	Exp int64 `json:"exp,omitempty"`

	// A key/value map that can contain custom parameters.
	Params map[string]string `json:"params,omitempty"`

	// The identifier for a resource to which this client is seeking access. The identifier MUST correspond to a resource that was previously registered.
	// Required: true
	ResourceID *string `json:"resource_id"`

	// An array referencing zero or more strings representing scopes to which access was granted for this resource. Each string MUST correspond to a scope that was registered by this resource server for the referenced resource.
	// Required: true
	ResourceScopes []string `json:"resource_scopes"`
}

// Validate validates this Uma permission list
func (m *UMAPermissionList) Validate(formats strfmt.Registry) error {
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

func (m *UMAPermissionList) validateResourceID(formats strfmt.Registry) error {

	if err := validate.Required("resource_id", "body", m.ResourceID); err != nil {
		return err
	}

	return nil
}

func (m *UMAPermissionList) validateResourceScopes(formats strfmt.Registry) error {

	if err := validate.Required("resource_scopes", "body", m.ResourceScopes); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UMAPermissionList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UMAPermissionList) UnmarshalBinary(b []byte) error {
	var res UMAPermissionList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
