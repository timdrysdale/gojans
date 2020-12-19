# IntrospectionResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Boolean indicator of whether or not the presented token is currently active. | [default to null]
**Scope** | **[]string** | Provide list of scopes to which access was granted for this resource. | [optional] [default to null]
**ClientId** | **string** | Client identifier for the OAuth 2.0 client that requested this token. | [optional] [default to null]
**Username** | **string** | Human-readable identifier for the resource owner who authorized this token. | [optional] [default to null]
**TokenType** | **string** | Type of the token as defined in Section 5.1 of OAuth 2.0 [RFC6749]. | [optional] [default to null]
**Exp** | **int32** | Integer timestamp, measured in the number of seconds since January 1 1970 UTC, indicating when this permission will expire. | [optional] [default to null]
**Iat** | **int32** |  | [optional] [default to null]
**Sub** | **string** | Subject of the token, as defined in JWT [RFC7519]. | [optional] [default to null]
**Aud** | **string** | Service-specific string identifier or list of string identifiers representing the intended audience for this token, as defined in JWT [RFC7519]. | [optional] [default to null]
**Iss** | **string** | String representing the issuer of this token, as defined in JWT [RFC7519]. | [optional] [default to null]
**AcrValues** | **string** | Authentication Context Class Reference values. | [optional] [default to null]
**Jti** | **string** | String identifier for the token, as defined in JWT. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

