// Code generated by go-swagger; DO NOT EDIT.

package fido2

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

// NewGetFIDO2ConfigurationParams creates a new GetFIDO2ConfigurationParams object
// with the default values initialized.
func NewGetFIDO2ConfigurationParams() *GetFIDO2ConfigurationParams {

	return &GetFIDO2ConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFIDO2ConfigurationParamsWithTimeout creates a new GetFIDO2ConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFIDO2ConfigurationParamsWithTimeout(timeout time.Duration) *GetFIDO2ConfigurationParams {

	return &GetFIDO2ConfigurationParams{

		timeout: timeout,
	}
}

// NewGetFIDO2ConfigurationParamsWithContext creates a new GetFIDO2ConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFIDO2ConfigurationParamsWithContext(ctx context.Context) *GetFIDO2ConfigurationParams {

	return &GetFIDO2ConfigurationParams{

		Context: ctx,
	}
}

// NewGetFIDO2ConfigurationParamsWithHTTPClient creates a new GetFIDO2ConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFIDO2ConfigurationParamsWithHTTPClient(client *http.Client) *GetFIDO2ConfigurationParams {

	return &GetFIDO2ConfigurationParams{
		HTTPClient: client,
	}
}

/*GetFIDO2ConfigurationParams contains all the parameters to send to the API endpoint
for the get fido2 configuration operation typically these are written to a http.Request
*/
type GetFIDO2ConfigurationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fido2 configuration params
func (o *GetFIDO2ConfigurationParams) WithTimeout(timeout time.Duration) *GetFIDO2ConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fido2 configuration params
func (o *GetFIDO2ConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fido2 configuration params
func (o *GetFIDO2ConfigurationParams) WithContext(ctx context.Context) *GetFIDO2ConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fido2 configuration params
func (o *GetFIDO2ConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fido2 configuration params
func (o *GetFIDO2ConfigurationParams) WithHTTPClient(client *http.Client) *GetFIDO2ConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fido2 configuration params
func (o *GetFIDO2ConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetFIDO2ConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}