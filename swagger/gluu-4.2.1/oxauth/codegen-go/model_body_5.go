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

type Body5 struct {
	// The token that the client wants to get revoked.
	Token string `json:"token"`
	// A hint about the type of the token submitted for revocation.
	TokenTypeHint string `json:"token_type_hint,omitempty"`
}
