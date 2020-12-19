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

// list of AuthenticateRequest
type RegisterRequestMessageAuthenticateRequests struct {
	// Version of the protocol that the to-be-registered U2F token must speak.
	Version string `json:"version,omitempty"`
	// The websafe-base64-encoded challenge.
	Challenge string `json:"challenge,omitempty"`
	// The application id that the RP would like to assert.
	AppId string `json:"appId,omitempty"`
	// websafe-base64 encoding of the key handle obtained from the U2F token during registration
	KeyHandle string `json:"keyHandle,omitempty"`
}
