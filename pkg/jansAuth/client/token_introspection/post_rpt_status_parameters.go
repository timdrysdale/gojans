// Code generated by go-swagger; DO NOT EDIT.

package token_introspection

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

// NewPostRptStatusParams creates a new PostRptStatusParams object
// with the default values initialized.
func NewPostRptStatusParams() *PostRptStatusParams {
	var ()
	return &PostRptStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRptStatusParamsWithTimeout creates a new PostRptStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRptStatusParamsWithTimeout(timeout time.Duration) *PostRptStatusParams {
	var ()
	return &PostRptStatusParams{

		timeout: timeout,
	}
}

// NewPostRptStatusParamsWithContext creates a new PostRptStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRptStatusParamsWithContext(ctx context.Context) *PostRptStatusParams {
	var ()
	return &PostRptStatusParams{

		Context: ctx,
	}
}

// NewPostRptStatusParamsWithHTTPClient creates a new PostRptStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRptStatusParamsWithHTTPClient(client *http.Client) *PostRptStatusParams {
	var ()
	return &PostRptStatusParams{
		HTTPClient: client,
	}
}

/*PostRptStatusParams contains all the parameters to send to the API endpoint
for the post rpt status operation typically these are written to a http.Request
*/
type PostRptStatusParams struct {

	/*Authorization
	  Client Authorization details that contains the access token along with other details.

	*/
	Authorization string
	/*Token
	  Client access token.

	*/
	Token string
	/*TokenTypeHint
	  ID Token previously issued by the Authorization Server being passed as a hint about the End-User.

	*/
	TokenTypeHint *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post rpt status params
func (o *PostRptStatusParams) WithTimeout(timeout time.Duration) *PostRptStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post rpt status params
func (o *PostRptStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post rpt status params
func (o *PostRptStatusParams) WithContext(ctx context.Context) *PostRptStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post rpt status params
func (o *PostRptStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post rpt status params
func (o *PostRptStatusParams) WithHTTPClient(client *http.Client) *PostRptStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post rpt status params
func (o *PostRptStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post rpt status params
func (o *PostRptStatusParams) WithAuthorization(authorization string) *PostRptStatusParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post rpt status params
func (o *PostRptStatusParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithToken adds the token to the post rpt status params
func (o *PostRptStatusParams) WithToken(token string) *PostRptStatusParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the post rpt status params
func (o *PostRptStatusParams) SetToken(token string) {
	o.Token = token
}

// WithTokenTypeHint adds the tokenTypeHint to the post rpt status params
func (o *PostRptStatusParams) WithTokenTypeHint(tokenTypeHint *string) *PostRptStatusParams {
	o.SetTokenTypeHint(tokenTypeHint)
	return o
}

// SetTokenTypeHint adds the tokenTypeHint to the post rpt status params
func (o *PostRptStatusParams) SetTokenTypeHint(tokenTypeHint *string) {
	o.TokenTypeHint = tokenTypeHint
}

// WriteToRequest writes these params to a swagger request
func (o *PostRptStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// form param token
	frToken := o.Token
	fToken := frToken
	if fToken != "" {
		if err := r.SetFormParam("token", fToken); err != nil {
			return err
		}
	}

	if o.TokenTypeHint != nil {

		// form param token_type_hint
		var frTokenTypeHint string
		if o.TokenTypeHint != nil {
			frTokenTypeHint = *o.TokenTypeHint
		}
		fTokenTypeHint := frTokenTypeHint
		if fTokenTypeHint != "" {
			if err := r.SetFormParam("token_type_hint", fTokenTypeHint); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
