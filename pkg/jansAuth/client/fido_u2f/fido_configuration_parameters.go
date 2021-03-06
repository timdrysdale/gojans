// Code generated by go-swagger; DO NOT EDIT.

package fido_u2f

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

// NewFIDOConfigurationParams creates a new FIDOConfigurationParams object
// with the default values initialized.
func NewFIDOConfigurationParams() *FIDOConfigurationParams {

	return &FIDOConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFIDOConfigurationParamsWithTimeout creates a new FIDOConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFIDOConfigurationParamsWithTimeout(timeout time.Duration) *FIDOConfigurationParams {

	return &FIDOConfigurationParams{

		timeout: timeout,
	}
}

// NewFIDOConfigurationParamsWithContext creates a new FIDOConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewFIDOConfigurationParamsWithContext(ctx context.Context) *FIDOConfigurationParams {

	return &FIDOConfigurationParams{

		Context: ctx,
	}
}

// NewFIDOConfigurationParamsWithHTTPClient creates a new FIDOConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFIDOConfigurationParamsWithHTTPClient(client *http.Client) *FIDOConfigurationParams {

	return &FIDOConfigurationParams{
		HTTPClient: client,
	}
}

/*FIDOConfigurationParams contains all the parameters to send to the API endpoint
for the fido configuration operation typically these are written to a http.Request
*/
type FIDOConfigurationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the fido configuration params
func (o *FIDOConfigurationParams) WithTimeout(timeout time.Duration) *FIDOConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fido configuration params
func (o *FIDOConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fido configuration params
func (o *FIDOConfigurationParams) WithContext(ctx context.Context) *FIDOConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fido configuration params
func (o *FIDOConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fido configuration params
func (o *FIDOConfigurationParams) WithHTTPClient(client *http.Client) *FIDOConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fido configuration params
func (o *FIDOConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FIDOConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
