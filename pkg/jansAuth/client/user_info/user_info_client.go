// Code generated by go-swagger; DO NOT EDIT.

package user_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user info API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user info API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetUserinfo(params *GetUserinfoParams) (*GetUserinfoOK, error)

	PostUserinfo(params *PostUserinfoParams) (*PostUserinfoOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetUserinfo gets userinfo

  Returns Claims about the authenticated End-User.
*/
func (a *Client) GetUserinfo(params *GetUserinfoParams) (*GetUserinfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserinfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-userinfo",
		Method:             "GET",
		PathPattern:        "/userinfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserinfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserinfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-userinfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostUserinfo posts userinfo

  Returns Claims about the authenticated End-User.
*/
func (a *Client) PostUserinfo(params *PostUserinfoParams) (*PostUserinfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostUserinfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "post-userinfo",
		Method:             "POST",
		PathPattern:        "/userinfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostUserinfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostUserinfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post-userinfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
