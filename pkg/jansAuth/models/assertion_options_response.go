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

// AssertionOptionsResponse AssertionOptionsResponse
//
// swagger:model AssertionOptionsResponse
type AssertionOptionsResponse struct {

	// allow credentials
	// Required: true
	AllowCredentials []string `json:"allowCredentials"`

	// Websafe-base64 encoding of the challenge.
	// Required: true
	Challenge *string `json:"challenge"`

	// extensions
	// Required: true
	Extensions *string `json:"extensions"`

	// username
	// Required: true
	User *string `json:"user"`

	// user verification
	// Required: true
	UserVerification *string `json:"userVerification"`
}

// Validate validates this assertion options response
func (m *AssertionOptionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChallenge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserVerification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssertionOptionsResponse) validateAllowCredentials(formats strfmt.Registry) error {

	if err := validate.Required("allowCredentials", "body", m.AllowCredentials); err != nil {
		return err
	}

	return nil
}

func (m *AssertionOptionsResponse) validateChallenge(formats strfmt.Registry) error {

	if err := validate.Required("challenge", "body", m.Challenge); err != nil {
		return err
	}

	return nil
}

func (m *AssertionOptionsResponse) validateExtensions(formats strfmt.Registry) error {

	if err := validate.Required("extensions", "body", m.Extensions); err != nil {
		return err
	}

	return nil
}

func (m *AssertionOptionsResponse) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	return nil
}

func (m *AssertionOptionsResponse) validateUserVerification(formats strfmt.Registry) error {

	if err := validate.Required("userVerification", "body", m.UserVerification); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssertionOptionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssertionOptionsResponse) UnmarshalBinary(b []byte) error {
	var res AssertionOptionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
