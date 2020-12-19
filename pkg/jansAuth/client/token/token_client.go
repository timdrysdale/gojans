// Code generated by go-swagger; DO NOT EDIT.

package token

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new token API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for token API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetIntrospection(params *GetIntrospectionParams) (*GetIntrospectionOK, error)

	PostIntrospection(params *PostIntrospectionParams) (*PostIntrospectionOK, error)

	PostToken(params *PostTokenParams) (*PostTokenOK, error)

	Revoke(params *RevokeParams) (*RevokeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetIntrospection gets introspection

  The Introspection OAuth 2 Endpoint.
*/
func (a *Client) GetIntrospection(params *GetIntrospectionParams) (*GetIntrospectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIntrospectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-introspection",
		Method:             "GET",
		PathPattern:        "/introspection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIntrospectionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIntrospectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-introspection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostIntrospection posts introspection

  The Introspection OAuth 2 Endpoint.
*/
func (a *Client) PostIntrospection(params *PostIntrospectionParams) (*PostIntrospectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostIntrospectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "post-introspection",
		Method:             "POST",
		PathPattern:        "/introspection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostIntrospectionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostIntrospectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post-introspection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostToken posts token

  To obtain an Access Token, an ID Token, and optionally a Refresh Token, the RP (Client).
*/
func (a *Client) PostToken(params *PostTokenParams) (*PostTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "post-token",
		Method:             "POST",
		PathPattern:        "/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostTokenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post-token: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Revoke revokes

  Revoke an Access Token or a Refresh Token, the RP (Client).
*/
func (a *Client) Revoke(params *RevokeParams) (*RevokeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRevokeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "revoke",
		Method:             "POST",
		PathPattern:        "/revoke",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RevokeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RevokeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for revoke: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}