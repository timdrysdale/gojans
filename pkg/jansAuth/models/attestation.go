// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Attestation Attestation
//
// list of fido2 attestation endpoints
//
// swagger:model Attestation
type Attestation struct {

	// fido2 attestation endpoint
	BasePath string `json:"base_path,omitempty"`

	// fido2 attestation options endpoint
	OptionsEnpoint string `json:"options_enpoint,omitempty"`

	// fido2 attestation result endpoint
	ResultEnpoint string `json:"result_enpoint,omitempty"`
}

// Validate validates this attestation
func (m *Attestation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Attestation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Attestation) UnmarshalBinary(b []byte) error {
	var res Attestation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}