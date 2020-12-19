// Code generated by go-swagger; DO NOT EDIT.

package fido2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new fido2 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for fido2 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AttestationOptions(params *AttestationOptionsParams) (*AttestationOptionsOK, error)

	AttestationResult(params *AttestationResultParams) (*AttestationResultOK, error)

	GetFIDO2Configuration(params *GetFIDO2ConfigurationParams) (*GetFIDO2ConfigurationOK, error)

	Options(params *OptionsParams) (*OptionsOK, error)

	Result(params *ResultParams) (*ResultOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AttestationOptions attestations options

  Created new registration.
*/
func (a *Client) AttestationOptions(params *AttestationOptionsParams) (*AttestationOptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttestationOptionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "attestation-options",
		Method:             "POST",
		PathPattern:        "/fido2/attestation/options",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AttestationOptionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AttestationOptionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for attestation-options: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AttestationResult attestations result

  FIDO2 attestation result
*/
func (a *Client) AttestationResult(params *AttestationResultParams) (*AttestationResultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttestationResultParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "attestation-result",
		Method:             "POST",
		PathPattern:        "/fido2/attestation/result",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AttestationResultReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AttestationResultOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for attestation-result: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFIDO2Configuration gets fido2 configuration

  FIDO2 configuration
*/
func (a *Client) GetFIDO2Configuration(params *GetFIDO2ConfigurationParams) (*GetFIDO2ConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFIDO2ConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-fido2-configuration",
		Method:             "GET",
		PathPattern:        "/fido2/configuration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFIDO2ConfigurationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFIDO2ConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-fido2-configuration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Options options

  FIDO2 Assertion Options
*/
func (a *Client) Options(params *OptionsParams) (*OptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOptionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "options",
		Method:             "POST",
		PathPattern:        "/fido2/assertion/options",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &OptionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*OptionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for options: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Result results

  FIDO2 Assertion Result.
*/
func (a *Client) Result(params *ResultParams) (*ResultOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResultParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "result",
		Method:             "POST",
		PathPattern:        "/fido2/assertion/result",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ResultReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ResultOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for result: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}