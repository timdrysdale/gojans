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

// GluuConfigurationResponse GluuConfigurationResponse
//
// Client GluuAttribute by Dn(Distinguished Name) based on Authorization Scope.
//
// swagger:model GluuConfigurationResponse
type GluuConfigurationResponse struct {

	// auth level mapping
	AuthLevelMapping map[string]string `json:"auth_level_mapping,omitempty"`

	// id generation endpoint
	// Required: true
	IDGenerationEndpoint *string `json:"id_generation_endpoint"`

	// introspection endpoint
	// Required: true
	IntrospectionEndpoint *string `json:"introspection_endpoint"`

	// scope to claims mapping
	ScopeToClaimsMapping map[string]string `json:"scope_to_claims_mapping,omitempty"`
}

// Validate validates this gluu configuration response
func (m *GluuConfigurationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIDGenerationEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntrospectionEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GluuConfigurationResponse) validateIDGenerationEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("id_generation_endpoint", "body", m.IDGenerationEndpoint); err != nil {
		return err
	}

	return nil
}

func (m *GluuConfigurationResponse) validateIntrospectionEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("introspection_endpoint", "body", m.IntrospectionEndpoint); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GluuConfigurationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GluuConfigurationResponse) UnmarshalBinary(b []byte) error {
	var res GluuConfigurationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
