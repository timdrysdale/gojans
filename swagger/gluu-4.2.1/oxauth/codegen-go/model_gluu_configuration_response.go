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

// Client GluuAttribute by Dn(Distinguished Name) based on Authorization Scope.
type GluuConfigurationResponse struct {
	IdGenerationEndpoint string `json:"id_generation_endpoint"`
	IntrospectionEndpoint string `json:"introspection_endpoint"`
	AuthLevelMapping map[string]string `json:"auth_level_mapping,omitempty"`
	ScopeToClaimsMapping map[string]string `json:"scope_to_claims_mapping,omitempty"`
}
