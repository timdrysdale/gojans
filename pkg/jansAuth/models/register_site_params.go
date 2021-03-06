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

// RegisterSiteParams RegisterSiteParams
//
// swagger:model RegisterSiteParams
type RegisterSiteParams struct {

	// token response
	TokenResponse *TokenResponse2 `json:"tokenResponse,omitempty"`

	// username
	// Required: true
	Username *string `json:"username"`
}

// Validate validates this register site params
func (m *RegisterSiteParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTokenResponse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegisterSiteParams) validateTokenResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.TokenResponse) { // not required
		return nil
	}

	if m.TokenResponse != nil {
		if err := m.TokenResponse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tokenResponse")
			}
			return err
		}
	}

	return nil
}

func (m *RegisterSiteParams) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegisterSiteParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegisterSiteParams) UnmarshalBinary(b []byte) error {
	var res RegisterSiteParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
