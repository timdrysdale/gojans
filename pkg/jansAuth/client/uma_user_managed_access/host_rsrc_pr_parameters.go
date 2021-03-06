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

// NewHostRsrcPrParams creates a new HostRsrcPrParams object
// with the default values initialized.
func NewHostRsrcPrParams() *HostRsrcPrParams {
	var ()
	return &HostRsrcPrParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHostRsrcPrParamsWithTimeout creates a new HostRsrcPrParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHostRsrcPrParamsWithTimeout(timeout time.Duration) *HostRsrcPrParams {
	var ()
	return &HostRsrcPrParams{

		timeout: timeout,
	}
}

// NewHostRsrcPrParamsWithContext creates a new HostRsrcPrParams object
// with the default values initialized, and the ability to set a context for a request
func NewHostRsrcPrParamsWithContext(ctx context.Context) *HostRsrcPrParams {
	var ()
	return &HostRsrcPrParams{

		Context: ctx,
	}
}

// NewHostRsrcPrParamsWithHTTPClient creates a new HostRsrcPrParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHostRsrcPrParamsWithHTTPClient(client *http.Client) *HostRsrcPrParams {
	var ()
	return &HostRsrcPrParams{
		HTTPClient: client,
	}
}

/*HostRsrcPrParams contains all the parameters to send to the API endpoint
for the host rsrc pr operation typically these are written to a http.Request
*/
type HostRsrcPrParams struct {

	/*Authorization
	  Client Authorization details that contains the access token along with other details.

	*/
	Authorization string
	/*Params
	  A key/value map that can contain custom parameters.

	*/
	Params *string
	/*ResourceID
	  The identifier for a resource to which this client is seeking access. The identifier MUST correspond to a resource that was previously registered.

	*/
	ResourceID string
	/*ResourceScopes
	  An array referencing zero or more strings representing scopes to which access was granted for this resource. Each string MUST correspond to a scope that was registered by this resource server for the referenced resource.

	*/
	ResourceScopes []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the host rsrc pr params
func (o *HostRsrcPrParams) WithTimeout(timeout time.Duration) *HostRsrcPrParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the host rsrc pr params
func (o *HostRsrcPrParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the host rsrc pr params
func (o *HostRsrcPrParams) WithContext(ctx context.Context) *HostRsrcPrParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the host rsrc pr params
func (o *HostRsrcPrParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the host rsrc pr params
func (o *HostRsrcPrParams) WithHTTPClient(client *http.Client) *HostRsrcPrParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the host rsrc pr params
func (o *HostRsrcPrParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the host rsrc pr params
func (o *HostRsrcPrParams) WithAuthorization(authorization string) *HostRsrcPrParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the host rsrc pr params
func (o *HostRsrcPrParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithParams adds the params to the host rsrc pr params
func (o *HostRsrcPrParams) WithParams(params *string) *HostRsrcPrParams {
	o.SetParams(params)
	return o
}

// SetParams adds the params to the host rsrc pr params
func (o *HostRsrcPrParams) SetParams(params *string) {
	o.Params = params
}

// WithResourceID adds the resourceID to the host rsrc pr params
func (o *HostRsrcPrParams) WithResourceID(resourceID string) *HostRsrcPrParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the host rsrc pr params
func (o *HostRsrcPrParams) SetResourceID(resourceID string) {
	o.ResourceID = resourceID
}

// WithResourceScopes adds the resourceScopes to the host rsrc pr params
func (o *HostRsrcPrParams) WithResourceScopes(resourceScopes []string) *HostRsrcPrParams {
	o.SetResourceScopes(resourceScopes)
	return o
}

// SetResourceScopes adds the resourceScopes to the host rsrc pr params
func (o *HostRsrcPrParams) SetResourceScopes(resourceScopes []string) {
	o.ResourceScopes = resourceScopes
}

// WriteToRequest writes these params to a swagger request
func (o *HostRsrcPrParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Params != nil {

		// form param params
		var frParams string
		if o.Params != nil {
			frParams = *o.Params
		}
		fParams := frParams
		if fParams != "" {
			if err := r.SetFormParam("params", fParams); err != nil {
				return err
			}
		}

	}

	// form param resource_id
	frResourceID := o.ResourceID
	fResourceID := frResourceID
	if fResourceID != "" {
		if err := r.SetFormParam("resource_id", fResourceID); err != nil {
			return err
		}
	}

	valuesResourceScopes := o.ResourceScopes

	joinedResourceScopes := swag.JoinByFormat(valuesResourceScopes, "")
	// form array param resource_scopes
	if err := r.SetFormParam("resource_scopes", joinedResourceScopes...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
