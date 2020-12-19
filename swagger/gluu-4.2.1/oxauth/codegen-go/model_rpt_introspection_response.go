/*
 * oxAuth
 *
 * oxAuth - OAuth 2.0 server; OpenID Connect Provider (OP) & UMA Authorization Server (AS)
 *
 * API version: 4.2
 * Contact: yuriyz@gluu.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RptIntrospectionResponse struct {
	// Boolean indicator of whether or not the presented token is currently active.
	Active bool `json:"active"`
	// Integer timestamp, in seconds since January 1 1970 UTC, indicating when this token will expire.
	Exp int64 `json:"exp,omitempty"`
	// Integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this permission was originally issued.
	Iat int32 `json:"iat,omitempty"`
	// Client id used to obtain RPT.
	ClientId string `json:"clientId,omitempty"`
	// Subject of the token. Usually a machine-readable identifier of the resource owner who authorized this token.
	Sub string `json:"sub,omitempty"`
	// Service-specific string identifier or list of string identifiers representing the intended audience for this token.
	Aud string `json:"aud,omitempty"`
	Permissions []RptIntrospectionResponsePermissions `json:"permissions"`
	// PCT token claims.
	PctClaims map[string]string `json:"pct_claims,omitempty"`
	// String representing the issuer of this token, as defined in JWT [RFC7519].
	Iss string `json:"iss,omitempty"`
	// String identifier for the token, as defined in JWT [RFC7519].
	Jti string `json:"jti,omitempty"`
	// Integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating the time before which this permission is not valid.
	Nbf int32 `json:"nbf,omitempty"`
	// Resource ID.
	ResourceId string `json:"resource_id"`
	ResourceScopes []string `json:"resource_scopes"`
}
