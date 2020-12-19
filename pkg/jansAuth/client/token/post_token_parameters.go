// Code generated by go-swagger; DO NOT EDIT.

package token

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

// NewPostTokenParams creates a new PostTokenParams object
// with the default values initialized.
func NewPostTokenParams() *PostTokenParams {
	var ()
	return &PostTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostTokenParamsWithTimeout creates a new PostTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostTokenParamsWithTimeout(timeout time.Duration) *PostTokenParams {
	var ()
	return &PostTokenParams{

		timeout: timeout,
	}
}

// NewPostTokenParamsWithContext creates a new PostTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostTokenParamsWithContext(ctx context.Context) *PostTokenParams {
	var ()
	return &PostTokenParams{

		Context: ctx,
	}
}

// NewPostTokenParamsWithHTTPClient creates a new PostTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostTokenParamsWithHTTPClient(client *http.Client) *PostTokenParams {
	var ()
	return &PostTokenParams{
		HTTPClient: client,
	}
}

/*PostTokenParams contains all the parameters to send to the API endpoint
for the post token operation typically these are written to a http.Request
*/
type PostTokenParams struct {

	/*Assertion
	  Assertion.

	*/
	Assertion *string
	/*ClaimToken*/
	ClaimToken *string
	/*ClaimTokenFormat*/
	ClaimTokenFormat *string
	/*ClientID
	  OAuth 2.0 Client Identifier valid at the Authorization Server.

	*/
	ClientID *string
	/*ClientSecret
	  The client secret.  The client MAY omit the parameter if the client secret is an empty string.

	*/
	ClientSecret *string
	/*Code
	  Code which is returned by authorization endpoint. (For grant_type=authorization_code)

	*/
	Code *string
	/*CodeVerifier
	  The client's PKCE code verifier.

	*/
	CodeVerifier *string
	/*GrantType
	  Provide a list of the OAuth 2.0 grant types that the Client is declaring that it will restrict itself to using.

	*/
	GrantType []string
	/*Password
	  End-User password.

	*/
	Password *string
	/*Pct*/
	Pct *string
	/*RedirectURI
	  Redirection URI to which the response will be sent. This URI MUST exactly match one of the Redirection URI values for the Client pre-registered at the OpenID Provider.

	*/
	RedirectURI *string
	/*RefreshToken
	  Refresh token.

	*/
	RefreshToken *string
	/*Rpt*/
	Rpt *string
	/*Scope
	  OpenID Connect requests MUST contain the openid scope value. If the openid scope value is not present, the behavior is entirely unspecified. Other scope values MAY be present. Scope values used that are not understood by an implementation SHOULD be ignored.

	*/
	Scope []string
	/*Ticket*/
	Ticket *string
	/*Username
	  End-User username.

	*/
	Username *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post token params
func (o *PostTokenParams) WithTimeout(timeout time.Duration) *PostTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post token params
func (o *PostTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post token params
func (o *PostTokenParams) WithContext(ctx context.Context) *PostTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post token params
func (o *PostTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post token params
func (o *PostTokenParams) WithHTTPClient(client *http.Client) *PostTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post token params
func (o *PostTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssertion adds the assertion to the post token params
func (o *PostTokenParams) WithAssertion(assertion *string) *PostTokenParams {
	o.SetAssertion(assertion)
	return o
}

// SetAssertion adds the assertion to the post token params
func (o *PostTokenParams) SetAssertion(assertion *string) {
	o.Assertion = assertion
}

// WithClaimToken adds the claimToken to the post token params
func (o *PostTokenParams) WithClaimToken(claimToken *string) *PostTokenParams {
	o.SetClaimToken(claimToken)
	return o
}

// SetClaimToken adds the claimToken to the post token params
func (o *PostTokenParams) SetClaimToken(claimToken *string) {
	o.ClaimToken = claimToken
}

// WithClaimTokenFormat adds the claimTokenFormat to the post token params
func (o *PostTokenParams) WithClaimTokenFormat(claimTokenFormat *string) *PostTokenParams {
	o.SetClaimTokenFormat(claimTokenFormat)
	return o
}

// SetClaimTokenFormat adds the claimTokenFormat to the post token params
func (o *PostTokenParams) SetClaimTokenFormat(claimTokenFormat *string) {
	o.ClaimTokenFormat = claimTokenFormat
}

// WithClientID adds the clientID to the post token params
func (o *PostTokenParams) WithClientID(clientID *string) *PostTokenParams {
	o.SetClientID(clientID)
	return o
}

// SetClientID adds the clientId to the post token params
func (o *PostTokenParams) SetClientID(clientID *string) {
	o.ClientID = clientID
}

// WithClientSecret adds the clientSecret to the post token params
func (o *PostTokenParams) WithClientSecret(clientSecret *string) *PostTokenParams {
	o.SetClientSecret(clientSecret)
	return o
}

// SetClientSecret adds the clientSecret to the post token params
func (o *PostTokenParams) SetClientSecret(clientSecret *string) {
	o.ClientSecret = clientSecret
}

// WithCode adds the code to the post token params
func (o *PostTokenParams) WithCode(code *string) *PostTokenParams {
	o.SetCode(code)
	return o
}

// SetCode adds the code to the post token params
func (o *PostTokenParams) SetCode(code *string) {
	o.Code = code
}

// WithCodeVerifier adds the codeVerifier to the post token params
func (o *PostTokenParams) WithCodeVerifier(codeVerifier *string) *PostTokenParams {
	o.SetCodeVerifier(codeVerifier)
	return o
}

// SetCodeVerifier adds the codeVerifier to the post token params
func (o *PostTokenParams) SetCodeVerifier(codeVerifier *string) {
	o.CodeVerifier = codeVerifier
}

// WithGrantType adds the grantType to the post token params
func (o *PostTokenParams) WithGrantType(grantType []string) *PostTokenParams {
	o.SetGrantType(grantType)
	return o
}

// SetGrantType adds the grantType to the post token params
func (o *PostTokenParams) SetGrantType(grantType []string) {
	o.GrantType = grantType
}

// WithPassword adds the password to the post token params
func (o *PostTokenParams) WithPassword(password *string) *PostTokenParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the post token params
func (o *PostTokenParams) SetPassword(password *string) {
	o.Password = password
}

// WithPct adds the pct to the post token params
func (o *PostTokenParams) WithPct(pct *string) *PostTokenParams {
	o.SetPct(pct)
	return o
}

// SetPct adds the pct to the post token params
func (o *PostTokenParams) SetPct(pct *string) {
	o.Pct = pct
}

// WithRedirectURI adds the redirectURI to the post token params
func (o *PostTokenParams) WithRedirectURI(redirectURI *string) *PostTokenParams {
	o.SetRedirectURI(redirectURI)
	return o
}

// SetRedirectURI adds the redirectUri to the post token params
func (o *PostTokenParams) SetRedirectURI(redirectURI *string) {
	o.RedirectURI = redirectURI
}

// WithRefreshToken adds the refreshToken to the post token params
func (o *PostTokenParams) WithRefreshToken(refreshToken *string) *PostTokenParams {
	o.SetRefreshToken(refreshToken)
	return o
}

// SetRefreshToken adds the refreshToken to the post token params
func (o *PostTokenParams) SetRefreshToken(refreshToken *string) {
	o.RefreshToken = refreshToken
}

// WithRpt adds the rpt to the post token params
func (o *PostTokenParams) WithRpt(rpt *string) *PostTokenParams {
	o.SetRpt(rpt)
	return o
}

// SetRpt adds the rpt to the post token params
func (o *PostTokenParams) SetRpt(rpt *string) {
	o.Rpt = rpt
}

// WithScope adds the scope to the post token params
func (o *PostTokenParams) WithScope(scope []string) *PostTokenParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the post token params
func (o *PostTokenParams) SetScope(scope []string) {
	o.Scope = scope
}

// WithTicket adds the ticket to the post token params
func (o *PostTokenParams) WithTicket(ticket *string) *PostTokenParams {
	o.SetTicket(ticket)
	return o
}

// SetTicket adds the ticket to the post token params
func (o *PostTokenParams) SetTicket(ticket *string) {
	o.Ticket = ticket
}

// WithUsername adds the username to the post token params
func (o *PostTokenParams) WithUsername(username *string) *PostTokenParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the post token params
func (o *PostTokenParams) SetUsername(username *string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *PostTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Assertion != nil {

		// form param assertion
		var frAssertion string
		if o.Assertion != nil {
			frAssertion = *o.Assertion
		}
		fAssertion := frAssertion
		if fAssertion != "" {
			if err := r.SetFormParam("assertion", fAssertion); err != nil {
				return err
			}
		}

	}

	if o.ClaimToken != nil {

		// form param claim_token
		var frClaimToken string
		if o.ClaimToken != nil {
			frClaimToken = *o.ClaimToken
		}
		fClaimToken := frClaimToken
		if fClaimToken != "" {
			if err := r.SetFormParam("claim_token", fClaimToken); err != nil {
				return err
			}
		}

	}

	if o.ClaimTokenFormat != nil {

		// form param claim_token_format
		var frClaimTokenFormat string
		if o.ClaimTokenFormat != nil {
			frClaimTokenFormat = *o.ClaimTokenFormat
		}
		fClaimTokenFormat := frClaimTokenFormat
		if fClaimTokenFormat != "" {
			if err := r.SetFormParam("claim_token_format", fClaimTokenFormat); err != nil {
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

	if o.ClientSecret != nil {

		// form param client_secret
		var frClientSecret string
		if o.ClientSecret != nil {
			frClientSecret = *o.ClientSecret
		}
		fClientSecret := frClientSecret
		if fClientSecret != "" {
			if err := r.SetFormParam("client_secret", fClientSecret); err != nil {
				return err
			}
		}

	}

	if o.Code != nil {

		// form param code
		var frCode string
		if o.Code != nil {
			frCode = *o.Code
		}
		fCode := frCode
		if fCode != "" {
			if err := r.SetFormParam("code", fCode); err != nil {
				return err
			}
		}

	}

	if o.CodeVerifier != nil {

		// form param code_verifier
		var frCodeVerifier string
		if o.CodeVerifier != nil {
			frCodeVerifier = *o.CodeVerifier
		}
		fCodeVerifier := frCodeVerifier
		if fCodeVerifier != "" {
			if err := r.SetFormParam("code_verifier", fCodeVerifier); err != nil {
				return err
			}
		}

	}

	valuesGrantType := o.GrantType

	joinedGrantType := swag.JoinByFormat(valuesGrantType, "")
	// form array param grant_type
	if err := r.SetFormParam("grant_type", joinedGrantType...); err != nil {
		return err
	}

	if o.Password != nil {

		// form param password
		var frPassword string
		if o.Password != nil {
			frPassword = *o.Password
		}
		fPassword := frPassword
		if fPassword != "" {
			if err := r.SetFormParam("password", fPassword); err != nil {
				return err
			}
		}

	}

	if o.Pct != nil {

		// form param pct
		var frPct string
		if o.Pct != nil {
			frPct = *o.Pct
		}
		fPct := frPct
		if fPct != "" {
			if err := r.SetFormParam("pct", fPct); err != nil {
				return err
			}
		}

	}

	if o.RedirectURI != nil {

		// form param redirect_uri
		var frRedirectURI string
		if o.RedirectURI != nil {
			frRedirectURI = *o.RedirectURI
		}
		fRedirectURI := frRedirectURI
		if fRedirectURI != "" {
			if err := r.SetFormParam("redirect_uri", fRedirectURI); err != nil {
				return err
			}
		}

	}

	if o.RefreshToken != nil {

		// form param refresh_token
		var frRefreshToken string
		if o.RefreshToken != nil {
			frRefreshToken = *o.RefreshToken
		}
		fRefreshToken := frRefreshToken
		if fRefreshToken != "" {
			if err := r.SetFormParam("refresh_token", fRefreshToken); err != nil {
				return err
			}
		}

	}

	if o.Rpt != nil {

		// form param rpt
		var frRpt string
		if o.Rpt != nil {
			frRpt = *o.Rpt
		}
		fRpt := frRpt
		if fRpt != "" {
			if err := r.SetFormParam("rpt", fRpt); err != nil {
				return err
			}
		}

	}

	valuesScope := o.Scope

	joinedScope := swag.JoinByFormat(valuesScope, "")
	// form array param scope
	if err := r.SetFormParam("scope", joinedScope...); err != nil {
		return err
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

	if o.Username != nil {

		// form param username
		var frUsername string
		if o.Username != nil {
			frUsername = *o.Username
		}
		fUsername := frUsername
		if fUsername != "" {
			if err := r.SetFormParam("username", fUsername); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}