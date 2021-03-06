// Code generated by go-swagger; DO NOT EDIT.

package registration

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

	"github.com/timdrysdale/gojans/pkg/jansAuth/models"
)

// NewPutRegisterParams creates a new PutRegisterParams object
// with the default values initialized.
func NewPutRegisterParams() *PutRegisterParams {
	var ()
	return &PutRegisterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutRegisterParamsWithTimeout creates a new PutRegisterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutRegisterParamsWithTimeout(timeout time.Duration) *PutRegisterParams {
	var ()
	return &PutRegisterParams{

		timeout: timeout,
	}
}

// NewPutRegisterParamsWithContext creates a new PutRegisterParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutRegisterParamsWithContext(ctx context.Context) *PutRegisterParams {
	var ()
	return &PutRegisterParams{

		Context: ctx,
	}
}

// NewPutRegisterParamsWithHTTPClient creates a new PutRegisterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutRegisterParamsWithHTTPClient(client *http.Client) *PutRegisterParams {
	var ()
	return &PutRegisterParams{
		HTTPClient: client,
	}
}

/*PutRegisterParams contains all the parameters to send to the API endpoint
for the put register operation typically these are written to a http.Request
*/
type PutRegisterParams struct {

	/*Authorization
	  Authorization header carrying \"registration_access_token\" issued before as a Bearer token

	*/
	Authorization string
	/*Body*/
	Body *models.RegisterParams
	/*ClientID
	  Client ID that identifies client that must be updated by this request.

	*/
	ClientID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put register params
func (o *PutRegisterParams) WithTimeout(timeout time.Duration) *PutRegisterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put register params
func (o *PutRegisterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put register params
func (o *PutRegisterParams) WithContext(ctx context.Context) *PutRegisterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put register params
func (o *PutRegisterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put register params
func (o *PutRegisterParams) WithHTTPClient(client *http.Client) *PutRegisterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put register params
func (o *PutRegisterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the put register params
func (o *PutRegisterParams) WithAuthorization(authorization string) *PutRegisterParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the put register params
func (o *PutRegisterParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithBody adds the body to the put register params
func (o *PutRegisterParams) WithBody(body *models.RegisterParams) *PutRegisterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put register params
func (o *PutRegisterParams) SetBody(body *models.RegisterParams) {
	o.Body = body
}

// WithClientID adds the clientID to the put register params
func (o *PutRegisterParams) WithClientID(clientID string) *PutRegisterParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the put register params
func (o *PutRegisterParams) SetClientID(clientID string) {
	o.ClientID = clientID
}

// WriteToRequest writes these params to a swagger request
func (o *PutRegisterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// query param client_id
	qrClientID := o.ClientID
	qClientID := qrClientID
	if qClientID != "" {
		if err := r.SetQueryParam("client_id", qClientID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
