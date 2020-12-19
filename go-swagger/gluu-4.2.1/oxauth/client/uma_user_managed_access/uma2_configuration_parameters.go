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
)

// NewUMA2ConfigurationParams creates a new UMA2ConfigurationParams object
// with the default values initialized.
func NewUMA2ConfigurationParams() *UMA2ConfigurationParams {

	return &UMA2ConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUMA2ConfigurationParamsWithTimeout creates a new UMA2ConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUMA2ConfigurationParamsWithTimeout(timeout time.Duration) *UMA2ConfigurationParams {

	return &UMA2ConfigurationParams{

		timeout: timeout,
	}
}

// NewUMA2ConfigurationParamsWithContext creates a new UMA2ConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewUMA2ConfigurationParamsWithContext(ctx context.Context) *UMA2ConfigurationParams {

	return &UMA2ConfigurationParams{

		Context: ctx,
	}
}

// NewUMA2ConfigurationParamsWithHTTPClient creates a new UMA2ConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUMA2ConfigurationParamsWithHTTPClient(client *http.Client) *UMA2ConfigurationParams {

	return &UMA2ConfigurationParams{
		HTTPClient: client,
	}
}

/*UMA2ConfigurationParams contains all the parameters to send to the API endpoint
for the uma2 configuration operation typically these are written to a http.Request
*/
type UMA2ConfigurationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the uma2 configuration params
func (o *UMA2ConfigurationParams) WithTimeout(timeout time.Duration) *UMA2ConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the uma2 configuration params
func (o *UMA2ConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the uma2 configuration params
func (o *UMA2ConfigurationParams) WithContext(ctx context.Context) *UMA2ConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the uma2 configuration params
func (o *UMA2ConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the uma2 configuration params
func (o *UMA2ConfigurationParams) WithHTTPClient(client *http.Client) *UMA2ConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the uma2 configuration params
func (o *UMA2ConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UMA2ConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
