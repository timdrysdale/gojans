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

type RegisterResponseParam struct {
	// Unique Client Identifier. It MUST NOT be currently valid for any other registered Client.
	ClientId string `json:"client_id"`
	// This value is used by Confidential Clients to authenticate to the Token Endpoint
	ClientSecret string `json:"client_secret,omitempty"`
	// Registration Access Token that can be used at the Client Configuration Endpoint to perform subsequent operations upon the Client registration.
	RegistrationAccessToken string `json:"registration_access_token,omitempty"`
	// Location of the Client Configuration Endpoint where the Registration Access Token can be used to perform subsequent operations upon the resulting Client registration.
	RegistrationClientUri string `json:"registration_client_uri,omitempty"`
	// Time at which the Client Identifier was issued.
	ClientIdIssuedAt int32 `json:"client_id_issued_at,omitempty"`
	// Time at which the client_secret will expire or 0 if it will not expire.
	ClientSecretExpiresAt int32 `json:"client_secret_expires_at,omitempty"`
}
