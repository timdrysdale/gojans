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

type InlineResponse201 struct {
	// The identifier for a resource to which this client is seeking access. The identifier MUST correspond to a resource that was previously registered.
	ResourceId string `json:"resource_id"`
	// An array referencing zero or more strings representing scopes to which access was granted for this resource. Each string MUST correspond to a scope that was registered by this resource server for the referenced resource.
	ResourceScopes []string `json:"resource_scopes"`
	// A key/value map that can contain custom parameters.
	Params map[string]string `json:"params,omitempty"`
	// Number of seconds since January 1 1970 UTC, indicating when this token will expire.
	Exp int64 `json:"exp,omitempty"`
}