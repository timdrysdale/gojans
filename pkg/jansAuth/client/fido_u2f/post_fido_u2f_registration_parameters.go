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

	"github.com/timdrysdale/gojans/pkg/jansAuth/models"
)

// NewPostFIDOU2FRegistrationParams creates a new PostFIDOU2FRegistrationParams object
// with the default values initialized.
func NewPostFIDOU2FRegistrationParams() *PostFIDOU2FRegistrationParams {
	var ()
	return &PostFIDOU2FRegistrationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostFIDOU2FRegistrationParamsWithTimeout creates a new PostFIDOU2FRegistrationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostFIDOU2FRegistrationParamsWithTimeout(timeout time.Duration) *PostFIDOU2FRegistrationParams {
	var ()
	return &PostFIDOU2FRegistrationParams{

		timeout: timeout,
	}
}

// NewPostFIDOU2FRegistrationParamsWithContext creates a new PostFIDOU2FRegistrationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostFIDOU2FRegistrationParamsWithContext(ctx context.Context) *PostFIDOU2FRegistrationParams {
	var ()
	return &PostFIDOU2FRegistrationParams{

		Context: ctx,
	}
}

// NewPostFIDOU2FRegistrationParamsWithHTTPClient creates a new PostFIDOU2FRegistrationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostFIDOU2FRegistrationParamsWithHTTPClient(client *http.Client) *PostFIDOU2FRegistrationParams {
	var ()
	return &PostFIDOU2FRegistrationParams{
		HTTPClient: client,
	}
}

/*PostFIDOU2FRegistrationParams contains all the parameters to send to the API endpoint
for the post fido u2f registration operation typically these are written to a http.Request
*/
type PostFIDOU2FRegistrationParams struct {

	/*Body*/
	Body *models.RegisterSiteParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post fido u2f registration params
func (o *PostFIDOU2FRegistrationParams) WithTimeout(timeout time.Duration) *PostFIDOU2FRegistrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post fido u2f registration params
func (o *PostFIDOU2FRegistrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post fido u2f registration params
func (o *PostFIDOU2FRegistrationParams) WithContext(ctx context.Context) *PostFIDOU2FRegistrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post fido u2f registration params
func (o *PostFIDOU2FRegistrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post fido u2f registration params
func (o *PostFIDOU2FRegistrationParams) WithHTTPClient(client *http.Client) *PostFIDOU2FRegistrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post fido u2f registration params
func (o *PostFIDOU2FRegistrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post fido u2f registration params
func (o *PostFIDOU2FRegistrationParams) WithBody(body *models.RegisterSiteParams) *PostFIDOU2FRegistrationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post fido u2f registration params
func (o *PostFIDOU2FRegistrationParams) SetBody(body *models.RegisterSiteParams) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostFIDOU2FRegistrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
