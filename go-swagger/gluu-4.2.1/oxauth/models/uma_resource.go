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

// UMAResource UmaResource
//
// Resource description
//
// swagger:model UmaResource
type UMAResource struct {

	// A human-readable string describing the resource
	Description string `json:"description,omitempty"`

	// number of seconds since January 1 1970 UTC, indicating when this token will expire.
	Exp int64 `json:"exp,omitempty"`

	// number of seconds since January 1 1970 UTC, indicating when the token was issued at
	Iat int64 `json:"iat,omitempty"`

	// A URI for a graphic icon representing the resource set. The referenced icon MAY be used by the authorization server in its resource owner user interface for the resource owner.
	IconURI string `json:"icon_uri,omitempty"`

	// A human-readable string describing a set of one or more resources. This name MAY be used by the authorization server in its resource owner user interface for the resource owner.
	Name string `json:"name,omitempty"`

	// An array of strings, any of which MAY be a URI, indicating the available scopes for this resource set. URIs MUST resolve to scope descriptions as defined in Section 2.1. Published scope descriptions MAY reside anywhere on the web; a resource server is not required to self-host scope descriptions and may wish to point to standardized scope descriptions residing elsewhere. It is the resource server's responsibility to ensure that scope description documents are accessible to authorization servers through GET calls to support any user interface requirements. The resource server and authorization server are presumed to have separately negotiated any required interpretation of scope handling not conveyed through scope descriptions.
	// Required: true
	ResourceScopes []string `json:"resource_scopes"`

	// scope expression
	ScopeExpression string `json:"scope_expression,omitempty"`

	// A string uniquely identifying the semantics of the resource set. For example, if the resource set consists of a single resource that is an identity claim that leverages standardized claim semantics for \"verified email address\", the value of this property could be an identifying URI for this claim.
	Type string `json:"type,omitempty"`
}

// Validate validates this Uma resource
func (m *UMAResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceScopes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UMAResource) validateResourceScopes(formats strfmt.Registry) error {

	if err := validate.Required("resource_scopes", "body", m.ResourceScopes); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UMAResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UMAResource) UnmarshalBinary(b []byte) error {
	var res UMAResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}