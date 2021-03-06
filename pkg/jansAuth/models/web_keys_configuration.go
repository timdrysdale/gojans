// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WebKeysConfiguration WebKeysConfiguration
//
// JSON Web Key Set - A JSON object that represents a set of JWKs. The JSON object MUST have a keys member, which is an array of JWKs.
//
// swagger:model WebKeysConfiguration
type WebKeysConfiguration struct {

	// List of JWK - A JSON object that represents a cryptographic key. The members of the object represent properties of the key, including its value.
	// Required: true
	Keys []*JSONWebKey `json:"keys"`
}

// Validate validates this web keys configuration
func (m *WebKeysConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebKeysConfiguration) validateKeys(formats strfmt.Registry) error {

	if err := validate.Required("keys", "body", m.Keys); err != nil {
		return err
	}

	for i := 0; i < len(m.Keys); i++ {
		if swag.IsZero(m.Keys[i]) { // not required
			continue
		}

		if m.Keys[i] != nil {
			if err := m.Keys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebKeysConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebKeysConfiguration) UnmarshalBinary(b []byte) error {
	var res WebKeysConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
