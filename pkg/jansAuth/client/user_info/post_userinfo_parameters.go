// Code generated by go-swagger; DO NOT EDIT.

package user_info

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

// NewPostUserinfoParams creates a new PostUserinfoParams object
// with the default values initialized.
func NewPostUserinfoParams() *PostUserinfoParams {
	var ()
	return &PostUserinfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUserinfoParamsWithTimeout creates a new PostUserinfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUserinfoParamsWithTimeout(timeout time.Duration) *PostUserinfoParams {
	var ()
	return &PostUserinfoParams{

		timeout: timeout,
	}
}

// NewPostUserinfoParamsWithContext creates a new PostUserinfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUserinfoParamsWithContext(ctx context.Context) *PostUserinfoParams {
	var ()
	return &PostUserinfoParams{

		Context: ctx,
	}
}

// NewPostUserinfoParamsWithHTTPClient creates a new PostUserinfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostUserinfoParamsWithHTTPClient(client *http.Client) *PostUserinfoParams {
	var ()
	return &PostUserinfoParams{
		HTTPClient: client,
	}
}

/*PostUserinfoParams contains all the parameters to send to the API endpoint
for the post userinfo operation typically these are written to a http.Request
*/
type PostUserinfoParams struct {

	/*Authorization
	  Client Authorization details that contains the access token along with other details.

	*/
	Authorization *string
	/*AccessToken
	  OAuth 2.0 Access Token.

	*/
	AccessToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post userinfo params
func (o *PostUserinfoParams) WithTimeout(timeout time.Duration) *PostUserinfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post userinfo params
func (o *PostUserinfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post userinfo params
func (o *PostUserinfoParams) WithContext(ctx context.Context) *PostUserinfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post userinfo params
func (o *PostUserinfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post userinfo params
func (o *PostUserinfoParams) WithHTTPClient(client *http.Client) *PostUserinfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post userinfo params
func (o *PostUserinfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post userinfo params
func (o *PostUserinfoParams) WithAuthorization(authorization *string) *PostUserinfoParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post userinfo params
func (o *PostUserinfoParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithAccessToken adds the accessToken to the post userinfo params
func (o *PostUserinfoParams) WithAccessToken(accessToken string) *PostUserinfoParams {
	o.SetAccessToken(accessToken)
	return o
}

// SetAccessToken adds the accessToken to the post userinfo params
func (o *PostUserinfoParams) SetAccessToken(accessToken string) {
	o.AccessToken = accessToken
}

// WriteToRequest writes these params to a swagger request
func (o *PostUserinfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Authorization != nil {

		// header param Authorization
		if err := r.SetHeaderParam("Authorization", *o.Authorization); err != nil {
			return err
		}

	}

	// form param access_token
	frAccessToken := o.AccessToken
	fAccessToken := frAccessToken
	if fAccessToken != "" {
		if err := r.SetFormParam("access_token", fAccessToken); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
