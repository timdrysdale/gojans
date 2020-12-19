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

// Client details in response.
type ClientInfoResponse struct {
	DisplayName string `json:"displayName,omitempty"`
	// XRI i-number
	Inum string `json:"inum,omitempty"`
	// oxAuth Appication type
	OxAuthAppType string `json:"oxAuthAppType,omitempty"`
	// oxAuth ID Token Signed Response Algorithm
	OxAuthIdTokenSignedResponseAlg string `json:"oxAuthIdTokenSignedResponseAlg,omitempty"`
	// Array of redirect URIs values used in the Authorization
	OxAuthRedirectURI []string `json:"oxAuthRedirectURI,omitempty"`
	// oxAuth Attribute Scope Id
	OxId string `json:"oxId,omitempty"`
	CustomAttributes []string `json:"custom_attributes,omitempty"`
}
