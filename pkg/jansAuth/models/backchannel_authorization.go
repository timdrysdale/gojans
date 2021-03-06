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

// BackchannelAuthorization BackchannelAuthorization
//
// swagger:model BackchannelAuthorization
type BackchannelAuthorization struct {

	// a u t h r e q ID
	// Required: true
	AUTHREQID *string `json:"AUTH_REQ_ID"`

	// e x p i r e s i n
	// Required: true
	EXPIRESIN *int32 `json:"EXPIRES_IN"`

	// i n t e r v a l
	// Required: true
	INTERVAL *int32 `json:"INTERVAL"`
}

// Validate validates this backchannel authorization
func (m *BackchannelAuthorization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAUTHREQID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEXPIRESIN(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateINTERVAL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackchannelAuthorization) validateAUTHREQID(formats strfmt.Registry) error {

	if err := validate.Required("AUTH_REQ_ID", "body", m.AUTHREQID); err != nil {
		return err
	}

	return nil
}

func (m *BackchannelAuthorization) validateEXPIRESIN(formats strfmt.Registry) error {

	if err := validate.Required("EXPIRES_IN", "body", m.EXPIRESIN); err != nil {
		return err
	}

	return nil
}

func (m *BackchannelAuthorization) validateINTERVAL(formats strfmt.Registry) error {

	if err := validate.Required("INTERVAL", "body", m.INTERVAL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackchannelAuthorization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackchannelAuthorization) UnmarshalBinary(b []byte) error {
	var res BackchannelAuthorization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
