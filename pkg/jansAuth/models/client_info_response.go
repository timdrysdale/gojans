// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ClientInfoResponse ClientInfoResponse
//
// Client details in response.
//
// swagger:model ClientInfoResponse
type ClientInfoResponse struct {

	// custom attributes
	CustomAttributes []string `json:"custom_attributes"`

	// display name
	DisplayName string `json:"displayName,omitempty"`

	// XRI i-number
	Inum string `json:"inum,omitempty"`

	// jansAuth Application type
	JansAuthAppType string `json:"jansAuthAppType,omitempty"`

	// jansAuth ID Token Signed Response Algorithm
	JansAuthIDTokenSignedResponseAlg string `json:"jansAuthIdTokenSignedResponseAlg,omitempty"`

	// Array of redirect URIs values used in the Authorization
	JansAuthRedirectURI []string `json:"jansAuthRedirectURI"`

	// jansAuth Attribute Scope Id
	OxID string `json:"oxId,omitempty"`
}

// Validate validates this client info response
func (m *ClientInfoResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClientInfoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClientInfoResponse) UnmarshalBinary(b []byte) error {
	var res ClientInfoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
