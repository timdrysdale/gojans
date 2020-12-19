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

// FIDO U2F device registration details
type Fidou2fregistrationTokenResponse struct {
	RegistrationData string `json:"registrationData,omitempty"`
	ClientData string `json:"clientData,omitempty"`
	DeviceData string `json:"deviceData,omitempty"`
}
