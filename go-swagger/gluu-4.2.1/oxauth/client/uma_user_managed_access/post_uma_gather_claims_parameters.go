// Code generated by go-swagger; DO NOT EDIT.

package uma_user_managed_access

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewPostUMAGatherClaimsParams creates a new PostUMAGatherClaimsParams object
// with the default values initialized.
func NewPostUMAGatherClaimsParams() *PostUMAGatherClaimsParams {
	var ()
	return &PostUMAGatherClaimsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUMAGatherClaimsParamsWithTimeout creates a new PostUMAGatherClaimsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUMAGatherClaimsParamsWithTimeout(timeout time.Duration) *PostUMAGatherClaimsParams {
	var ()
	return &PostUMAGatherClaimsParams{

		timeout: timeout,
	}
}

// NewPostUMAGatherClaimsParamsWithContext creates a new PostUMAGatherClaimsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUMAGatherClaimsParamsWithContext(ctx context.Context) *PostUMAGatherClaimsParams {
	var ()
	return &PostUMAGatherClaimsParams{

		Context: ctx,
	}
}

// NewPostUMAGatherClaimsParamsWithHTTPClient creates a new PostUMAGatherClaimsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostUMAGatherClaimsParamsWithHTTPClient(client *http.Client) *PostUMAGatherClaimsParams {
	var ()
	return &PostUMAGatherClaimsParams{
		HTTPClient: client,
	}
}

/*PostUMAGatherClaimsParams contains all the parameters to send to the API endpoint
for the post uma gather claims operation typically these are written to a http.Request
*/
type PostUMAGatherClaimsParams struct {

	/*Authentication*/
	Authentication *bool
	/*ClaimsRedirectURI*/
	ClaimsRedirectURI *string
	/*ClientID
	  OAuth 2.0 Client Identifier valid at the Authorization Server.

	*/
	ClientID *string
	/*Reset*/
	Reset *bool
	/*State*/
	State *string
	/*Ticket*/
	Ticket *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post uma gather claims params
func (o *PostUMAGatherClaimsParams) WithTimeout(timeout time.Duration) *PostUMAGatherClaimsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post uma gather claims params
func (o *PostUMAGatherClaimsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post uma gather claims params
func (o *PostUMAGatherClaimsParams) WithContext(ctx context.Context) *PostUMAGatherClaimsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post uma gather claims params
func (o *PostUMAGatherClaimsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post uma gather claims params
func (o *PostUMAGatherClaimsParams) WithHTTPClient(client *http.Client) *PostUMAGatherClaimsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post uma gather claims params
func (o *PostUMAGatherClaimsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthentication adds the authentication to the post uma gather claims params
func (o *PostUMAGatherClaimsParams) WithAuthentication(authentication *bool) *PostUMAGatherClaimsParams {
	o.SetAuthentication(authentication)
	return o
}

// SetAuthentication adds the authentication to the post uma gather claims params
func (o *PostUMAGatherClaimsParams) SetAuthentication(authentication *bool) {
	o.Authentication = authentication
}

// WithClaimsRedirectURI adds the claimsRedirectURI to the post uma gather claims params
func (o *PostUMAGatherClaimsParams) WithClaimsRedirectURI(claimsRedirectURI *string) *PostUMAGatherClaimsParams {
	o.SetClaimsRedirectURI(claimsRedirectURI)
	return o
}

// SetClaimsRedirectURI adds the claimsRedirectUri to the post uma gather claims params
func (o *PostUMAGatherClaimsParams) SetClaimsRedirectURI(claimsRedirectURI *string) {
	o.ClaimsRedirectURI = claimsRedirectURI
}

// WithClientID adds the clientID to the post uma gather claims params
func (o *PostUMAGatherClaimsParams) WithClientID(clientID *string) *PostUMAGatherClaimsParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the post uma gather claims params
func (o *PostUMAGatherClaimsParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithReset adds the reset to the post uma gather claims params
func (o *PostUMAGatherClaimsParams) WithReset(reset *bool) *PostUMAGatherClaimsParams {
	o.SetReset(reset)
	return o
}

// SetReset adds the reset to the post uma gather claims params
func (o *PostUMAGatherClaimsParams) SetReset(reset *bool) {
	o.Reset = reset
}

// WithState adds the state to the post uma gather claims params
func (o *PostUMAGatherClaimsParams) WithState(state *string) *PostUMAGatherClaimsParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the post uma gather claims params
func (o *PostUMAGatherClaimsParams) SetState(state *string) {
	o.State = state
}

// WithTicket adds the ticket to the post uma gather claims params
func (o *PostUMAGatherClaimsParams) WithTicket(ticket *string) *PostUMAGatherClaimsParams {
	o.SetTicket(ticket)
	return o
}

// SetTicket adds the ticket to the post uma gather claims params
func (o *PostUMAGatherClaimsParams) SetTicket(ticket *string) {
	o.Ticket = ticket
}

// WriteToRequest writes these params to a swagger request
func (o *PostUMAGatherClaimsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Authentication != nil {

		// form param authentication
		var frAuthentication bool
		if o.Authentication != nil {
			frAuthentication = *o.Authentication
		}
		fAuthentication := swag.FormatBool(frAuthentication)
		if fAuthentication != "" {
			if err := r.SetFormParam("authentication", fAuthentication); err != nil {
				return err
			}
		}

	}

	if o.ClaimsRedirectURI != nil {

		// form param claims_redirect_uri
		var frClaimsRedirectURI string
		if o.ClaimsRedirectURI != nil {
			frClaimsRedirectURI = *o.ClaimsRedirectURI
		}
		fClaimsRedirectURI := frClaimsRedirectURI
		if fClaimsRedirectURI != "" {
			if err := r.SetFormParam("claims_redirect_uri", fClaimsRedirectURI); err != nil {
				return err
			}
		}

	}

	if o.ClientID != nil {

		// form param client_id
		var frClientID string
		if o.ClientID != nil {
			frClientID = *o.ClientID
		}
		fClientID := frClientID
		if fClientID != "" {
			if err := r.SetFormParam("client_id", fClientID); err != nil {
				return err
			}
		}

	}

	if o.Reset != nil {

		// form param reset
		var frReset bool
		if o.Reset != nil {
			frReset = *o.Reset
		}
		fReset := swag.FormatBool(frReset)
		if fReset != "" {
			if err := r.SetFormParam("reset", fReset); err != nil {
				return err
			}
		}

	}

	if o.State != nil {

		// form param state
		var frState string
		if o.State != nil {
			frState = *o.State
		}
		fState := frState
		if fState != "" {
			if err := r.SetFormParam("state", fState); err != nil {
				return err
			}
		}

	}

	if o.Ticket != nil {

		// form param ticket
		var frTicket string
		if o.Ticket != nil {
			frTicket = *o.Ticket
		}
		fTicket := frTicket
		if fTicket != "" {
			if err := r.SetFormParam("ticket", fTicket); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
