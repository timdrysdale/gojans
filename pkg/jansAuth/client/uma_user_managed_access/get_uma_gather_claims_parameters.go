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

// NewGetUMAGatherClaimsParams creates a new GetUMAGatherClaimsParams object
// with the default values initialized.
func NewGetUMAGatherClaimsParams() *GetUMAGatherClaimsParams {
	var ()
	return &GetUMAGatherClaimsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUMAGatherClaimsParamsWithTimeout creates a new GetUMAGatherClaimsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUMAGatherClaimsParamsWithTimeout(timeout time.Duration) *GetUMAGatherClaimsParams {
	var ()
	return &GetUMAGatherClaimsParams{

		timeout: timeout,
	}
}

// NewGetUMAGatherClaimsParamsWithContext creates a new GetUMAGatherClaimsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUMAGatherClaimsParamsWithContext(ctx context.Context) *GetUMAGatherClaimsParams {
	var ()
	return &GetUMAGatherClaimsParams{

		Context: ctx,
	}
}

// NewGetUMAGatherClaimsParamsWithHTTPClient creates a new GetUMAGatherClaimsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUMAGatherClaimsParamsWithHTTPClient(client *http.Client) *GetUMAGatherClaimsParams {
	var ()
	return &GetUMAGatherClaimsParams{
		HTTPClient: client,
	}
}

/*GetUMAGatherClaimsParams contains all the parameters to send to the API endpoint
for the get uma gather claims operation typically these are written to a http.Request
*/
type GetUMAGatherClaimsParams struct {

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

// WithTimeout adds the timeout to the get uma gather claims params
func (o *GetUMAGatherClaimsParams) WithTimeout(timeout time.Duration) *GetUMAGatherClaimsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get uma gather claims params
func (o *GetUMAGatherClaimsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get uma gather claims params
func (o *GetUMAGatherClaimsParams) WithContext(ctx context.Context) *GetUMAGatherClaimsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get uma gather claims params
func (o *GetUMAGatherClaimsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get uma gather claims params
func (o *GetUMAGatherClaimsParams) WithHTTPClient(client *http.Client) *GetUMAGatherClaimsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get uma gather claims params
func (o *GetUMAGatherClaimsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthentication adds the authentication to the get uma gather claims params
func (o *GetUMAGatherClaimsParams) WithAuthentication(authentication *bool) *GetUMAGatherClaimsParams {
	o.SetAuthentication(authentication)
	return o
}

// SetAuthentication adds the authentication to the get uma gather claims params
func (o *GetUMAGatherClaimsParams) SetAuthentication(authentication *bool) {
	o.Authentication = authentication
}

// WithClaimsRedirectURI adds the claimsRedirectURI to the get uma gather claims params
func (o *GetUMAGatherClaimsParams) WithClaimsRedirectURI(claimsRedirectURI *string) *GetUMAGatherClaimsParams {
	o.SetClaimsRedirectURI(claimsRedirectURI)
	return o
}

// SetClaimsRedirectURI adds the claimsRedirectUri to the get uma gather claims params
func (o *GetUMAGatherClaimsParams) SetClaimsRedirectURI(claimsRedirectURI *string) {
	o.ClaimsRedirectURI = claimsRedirectURI
}

// WithClientID adds the clientID to the get uma gather claims params
func (o *GetUMAGatherClaimsParams) WithClientID(clientID *string) *GetUMAGatherClaimsParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the get uma gather claims params
func (o *GetUMAGatherClaimsParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithReset adds the reset to the get uma gather claims params
func (o *GetUMAGatherClaimsParams) WithReset(reset *bool) *GetUMAGatherClaimsParams {
	o.SetReset(reset)
	return o
}

// SetReset adds the reset to the get uma gather claims params
func (o *GetUMAGatherClaimsParams) SetReset(reset *bool) {
	o.Reset = reset
}

// WithState adds the state to the get uma gather claims params
func (o *GetUMAGatherClaimsParams) WithState(state *string) *GetUMAGatherClaimsParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the get uma gather claims params
func (o *GetUMAGatherClaimsParams) SetState(state *string) {
	o.State = state
}

// WithTicket adds the ticket to the get uma gather claims params
func (o *GetUMAGatherClaimsParams) WithTicket(ticket *string) *GetUMAGatherClaimsParams {
	o.SetTicket(ticket)
	return o
}

// SetTicket adds the ticket to the get uma gather claims params
func (o *GetUMAGatherClaimsParams) SetTicket(ticket *string) {
	o.Ticket = ticket
}

// WriteToRequest writes these params to a swagger request
func (o *GetUMAGatherClaimsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Authentication != nil {

		// query param authentication
		var qrAuthentication bool
		if o.Authentication != nil {
			qrAuthentication = *o.Authentication
		}
		qAuthentication := swag.FormatBool(qrAuthentication)
		if qAuthentication != "" {
			if err := r.SetQueryParam("authentication", qAuthentication); err != nil {
				return err
			}
		}

	}

	if o.ClaimsRedirectURI != nil {

		// query param claims_redirect_uri
		var qrClaimsRedirectURI string
		if o.ClaimsRedirectURI != nil {
			qrClaimsRedirectURI = *o.ClaimsRedirectURI
		}
		qClaimsRedirectURI := qrClaimsRedirectURI
		if qClaimsRedirectURI != "" {
			if err := r.SetQueryParam("claims_redirect_uri", qClaimsRedirectURI); err != nil {
				return err
			}
		}

	}

	if o.ClientID != nil {

		// query param client_id
		var qrClientID string
		if o.ClientID != nil {
			qrClientID = *o.ClientID
		}
		qClientID := qrClientID
		if qClientID != "" {
			if err := r.SetQueryParam("client_id", qClientID); err != nil {
				return err
			}
		}

	}

	if o.Reset != nil {

		// query param reset
		var qrReset bool
		if o.Reset != nil {
			qrReset = *o.Reset
		}
		qReset := swag.FormatBool(qrReset)
		if qReset != "" {
			if err := r.SetQueryParam("reset", qReset); err != nil {
				return err
			}
		}

	}

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}

	}

	if o.Ticket != nil {

		// query param ticket
		var qrTicket string
		if o.Ticket != nil {
			qrTicket = *o.Ticket
		}
		qTicket := qrTicket
		if qTicket != "" {
			if err := r.SetQueryParam("ticket", qTicket); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}