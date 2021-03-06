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

// RegisterResponseParam RegisterResponseParam
//
// swagger:model RegisterResponseParam
type RegisterResponseParam struct {

	// Unique Client Identifier. It MUST NOT be currently valid for any other registered Client.
	// Required: true
	ClientID *string `json:"client_id"`

	// Time at which the Client Identifier was issued.
	ClientIDIssuedAt int32 `json:"client_id_issued_at,omitempty"`

	// This value is used by Confidential Clients to authenticate to the Token Endpoint
	ClientSecret string `json:"client_secret,omitempty"`

	// Time at which the client_secret will expire or 0 if it will not expire.
	ClientSecretExpiresAt int32 `json:"client_secret_expires_at,omitempty"`

	// Registration Access Token that can be used at the Client Configuration Endpoint to perform subsequent operations upon the Client registration.
	RegistrationAccessToken string `json:"registration_access_token,omitempty"`

	// Location of the Client Configuration Endpoint where the Registration Access Token can be used to perform subsequent operations upon the resulting Client registration.
	RegistrationClientURI string `json:"registration_client_uri,omitempty"`
}

// Validate validates this register response param
func (m *RegisterResponseParam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegisterResponseParam) validateClientID(formats strfmt.Registry) error {

	if err := validate.Required("client_id", "body", m.ClientID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegisterResponseParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegisterResponseParam) UnmarshalBinary(b []byte) error {
	var res RegisterResponseParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
